#!/bin/bash

if [[ -z "${SAMPLE_USER_DATA}" ]]; then
    export SAMPLE_USER_DATA="<not-set-path/to/user_data.json>"
fi

if [[ -z "${GIT_COMMIT}" ]]; then
    export GIT_COMMIT=$(git rev-parse HEAD)
fi

if [[ -z "${TAG_DATE}" ]]; then
    export TAG_DATE=$(date)
fi

if [[ -z "${BRANCH_NAME}" ]]; then
    export BRANCH_NAME=$(git branch --show-current)
fi

if [[ -z "${GIT_BRANCH}" ]]; then
    export GIT_BRANCH="$BRANCH_NAME@$GIT_COMMIT" 
fi

if [[ -z "${DB_CONNECTIONSTRING}" ]]; then
    export DB_CONNECTIONSTRING="postgres://user:pass@localhost/postgres?sslmode=disable"
fi