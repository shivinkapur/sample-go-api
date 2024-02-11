#!/bin/sh

# capture current date/time in UTC
unset $BUILD_TIMESTAMP
BUILD_TIMESTAMP=$(date -u +'%Y%m%d%H%M%S')
export BUILD_TIMESTAMP=$(echo $BUILD_TIMESTAMP)
echo 'Build Timestamp:' $BUILD_TIMESTAMP

# capture current commit hash
unset $GITHUB_SHA
GITHUB_SHA=$(git log -1 --format="%H")
export GITHUB_SHA=$(echo $GITHUB_SHA)
echo 'Commit:' $GITHUB_SHA

# capture current branch name
unset $GITHUB_REF_NAME
GITHUB_REF_NAME=$(git symbolic-ref --short HEAD)
export GITHUB_REF_NAME=$(echo $GITHUB_REF_NAME)
echo 'Branch:' $GITHUB_REF_NAME

# capture root
unset $PROJECT_ROOT
PROJECT_ROOT=$(pwd)
echo 'Project Root:' $PROJECT_ROOT

# create build output folder
OUTPUT_DIR_REF=$PROJECT_ROOT/target
echo "creating output directory:" $OUTPUT_DIR_REF
mkdir -p $OUTPUT_DIR_REF

# create output config folder for api config files
echo "creating output config directory for api:" $OUTPUT_DIR_REF/config/api
mkdir -p $OUTPUT_DIR_REF/config/api

# create temporary file
TEMP_API_VERSION=$(mktemp)

# copy contents of version info into temp file
cat $PROJECT_ROOT/api/version.json > $TEMP_API_VERSION
echo 'Copied contents of' $PROJECT_ROOT/api/version.json to $TEMP_API_VERSION

# replace values from env
envsubst < $TEMP_API_VERSION > $PROJECT_ROOT/api/version.json
echo 'Incorporating API Version Info:'
cat $PROJECT_ROOT/api/version.json

# build Fixed Income Services
echo 'Building Sample go lang API Services...'
if go clean -cache && go build -v -o $OUTPUT_DIR_REF/ ./... ; then
  echo 'Build succeeded'
  echo 'Copying api config files to'$OUTPUT_DIR_REF/config/api
  cp $PROJECT_ROOT/config/user_data.json $OUTPUT_DIR_REF/config/api/
else
  echo 'Build failed'
fi

# copy back unmodified version info for API
cat $TEMP_API_VERSION > $PROJECT_ROOT/api/version.json
echo 'Copied contents of' $TEMP_API_VERSION 'to' $PROJECT_ROOT/api/version.json

# delete temporary version copy for API
rm $TEMP_API_VERSION
echo 'Removed' $TEMP_API_VERSION

# Path: build.sh