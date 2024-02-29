#!/usr/bin/env bash

set -evo pipefail

ROOT="$(pwd)"
VERSION=$(sed -n '/version/ s/.*"\(.*\)".*/\1/p' < dist/manifest.yaml)

# LATEST_TAG=$(git describe --tags --abbrev=0 || git rev-parse --abbrev-ref HEAD)

make_pr_to_cli_index() {
    configure_git

    REPO='git@github.com:hasura/cli-plugins-index.git'
    echo "updating hasura/cli-plugins-index"

    export DIST_PATH="${ROOT}/dist"

    configure_git

    git clone ${REPO} ~/plugins-index
   
    cd ~/plugins-index
    # git checkout -b upgrade-from-v2-${VERSION}
    mkdir -p ./plugins/upgrade-from-v2/${VERSION}
    cp ${DIST_PATH}/manifest.yaml ./plugins/upgrade-from-v2/${VERSION}/manifest.yaml
    
    git add .
    git commit -m "update pro manifest to ${VERSION}"
    git push origin master # upgrade-from-v2-${VERSION}
}

configure_git() {
  git config --global user.email "build@hasura.io"
  git config --global user.name "hasura-bot"
}

make_pr_to_cli_index
