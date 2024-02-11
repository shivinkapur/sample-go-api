#!/bin/sh

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

# build Fixed Income Services
echo 'Building Sample go lang API Services...'
if go clean -cache && go build -v -o $OUTPUT_DIR_REF/ ./... ; then
  echo 'Build succeeded'
  echo 'Copying api config files to'$OUTPUT_DIR_REF/config/api
  cp $PROJECT_ROOT/config/user_data.json $OUTPUT_DIR_REF/config/api/
else
  echo 'Build failed'
fi

# Path: build.sh