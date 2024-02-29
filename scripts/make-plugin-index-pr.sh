#!/usr/bin/env bash

set -evo pipefail

ROOT="$(pwd)"
LATEST_TAG=$(git describe --tags --abbrev=0 || git rev-parse --abbrev-ref HEAD)

make_pr_to_cli_index() {
    configure_git

    REPO='git@github.com:hasura/cli-plugins-index.git'
    echo "updating hasura/cli-plugins-index"

    export DIST_PATH="${ROOT}/dist"

    configure_ssh
    configure_git

    git clone ${REPO} ~/plugins-index
   
    cd ~/plugins-index
    # TODO: Switch to main after we check that this works
    git checkout -b upgrade-from-v2-${LATEST_TAG}
    mkdir -p ./plugins/upgrade-from-v2/${LATEST_TAG}
    cp ${DIST_PATH}/manifest.yaml ./plugins/upgrade-from-v2/${LATEST_TAG}/manifest.yaml
    
    git add .
    git commit -m "update pro manifest to ${LATEST_TAG}"
    git push upgrade-from-v2-${LATEST_TAG}
}

configure_ssh() {
  mkdir -p ~/.ssh
  touch ~/.ssh/config
  echo "${INDEX_PRIVATE_KEY}" > ~/.ssh/cli-plugins-index
  echo "${INDEX_PUBLIC_KEY}" > ~/.ssh/cli-plugins-index.pub
  # echo "" >> ~/.ssh/cli-plugins-index
  # echo "# github.com/hasura/cli-plugins-index.git" >> ~/.ssh/cli-plugins-index
  chmod 600 ~/.ssh/cli-plugins-index
  cat <<EOF > ~/.ssh/config
  Host github.com-cli-plugins-index
	 HostName github.com
   User git
	 IdentityFile ~/.ssh/cli-plugins-index
EOF

  # Check config
  cat ~/.ssh/cli-plugins-index | sed 's/./X/g'
}

configure_git() {
  git config --global user.email "build@hasura.io"
  git config --global user.name "hasura-bot"
}

make_pr_to_cli_index
