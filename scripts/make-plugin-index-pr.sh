#!/usr/bin/env bash

set -evo pipefail

ROOT="$(pwd)"
LATEST_TAG=$(git describe --tags --abbrev=0)

make_pr_to_cli_index() {
    configure_git

    REPO='github.com/hasura/cli-plugins-index.git'
    echo "sending pr to hasura/cli-plugins-index"

    export DIST_PATH="${ROOT}/dist"

    configure_git
    git clone https://${REPO} ~/plugins-index
   
    cd ~/plugins-index
    git checkout -b upgrade-from-v2-${LATEST_TAG}
    mkdir -p ./plugins/upgrade-from-v2/${LATEST_TAG}
    cp ${DIST_PATH}/manifest.yaml ./plugins/upgrade-from-v2/${LATEST_TAG}/manifest.yaml
    
    git add .
    git commit -m "update pro manifest to ${LATEST_TAG}"
    git push -q https://${GITHUB_TOKEN}@${REPO} upgrade-from-v2-${LATEST_TAG}
    # hub pull-request -f -F- <<<"update upgrade-from-v2 manifest to ${LATEST_TAG}" -r ${REVIEWERS} -a ${REVIEWERS}
    hub pull-request -f -F- <<<"update upgrade-from-v2 manifest to ${LATEST_TAG}"

    unset DIST_PATH
}

configure_git() {
  git config --global user.email "build@hasura.io"
  git config --global user.name "hasura-bot"
}

make_pr_to_cli_index