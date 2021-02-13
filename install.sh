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
rm -rf .git go.mod go.sum install.sh
sed -i "s/go-boilerplate/${project_name}/g" main.go
sed -i "s/go-boilerplate/${project_name}/g" **/*.go
mv .go-boilerplate.yml ".${project_name}.yml"
git init && git remote add origin gh:l-lin/${project_name}
go mod init ${project_module}
cat > README.md <<EOF
# ${project_name}

![Go](https://github.com/l-lin/${project_name}/workflows/Go/badge.svg)

> Project's description

## Getting started

\`\`\`bash
# Build
make compile
\`\`\`

## Usage

\`\`\`bash
# Run binary
./bin/${project_name} -h
# Or directly using go
go run .
\`\`\`
EOF

echo "Project ${project_name} successfully initialized!"

} # this ensures the entire script is downloaded #
