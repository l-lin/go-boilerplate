#!/usr/bin/env bash

{ # this ensures the entire script is downloaded #

set -e

url=https://github.com/l-lin/go-boilerplate.git

if type git >/dev/null 2>&1; then
    if type go >/dev/null 2>&1; then
        echo "Preconditions are fulfilled. Proceeding..."
    else
        echo "Please install Go"
        exit 1
    fi
else
    echo "Please install Git"
    exit 1
fi

project_name=$(basename "$(pwd)")
project_module="github.com/l-lin/${project_name}"

git clone ${url} .
rm -rf .git go.mod install.sh
git init && git remote add origin https://${project_module}
go mod init ${project_module}
echo "# ${project_name}" > README.md

echo "Project ${project_name} successfully initialized!"

} # this ensures the entire script is downloaded #
