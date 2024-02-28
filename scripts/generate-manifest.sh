#!/usr/bin/env bash

set -evo pipefail
ROOT="$(pwd)"

export VERSION=$(${ROOT}/scripts/get-version.sh)
export BUCKET_URL=https://storage.googleapis.com/TODO/${VERSION}
export LINUX_AMD64_SHA256=$(cat ${ROOT}/dist/*.sha256   | grep linux-amd64   | cut -f1 -d' ')
export MACOS_AMD64_SHA256=$(cat ${ROOT}/dist/*.sha256   | grep darwin-amd64  | cut -f1 -d' ')
export WINDOWS_AMD64_SHA256=$(cat ${ROOT}/dist/*.sha256 | grep windows-amd64 | cut -f1 -d' ')
export LINUX_ARM64_SHA256=$(cat ${ROOT}/dist/*.sha256   | grep linux-arm64   | cut -f1 -d' ')
export MACOS_ARM64_SHA256=$(cat ${ROOT}/dist/*.sha256   | grep darwin-arm64  | cut -f1 -d' ')
export WINDOWS_ARM64_SHA256=$(cat ${ROOT}/dist/*.sha256 | grep windows-arm64 | cut -f1 -d' ')

(echo "cat <<EOF >${ROOT}/dist/manifest.yaml";
cat ${ROOT}/scripts/manifest.yaml; echo; echo EOF;
)>${ROOT}/dist/manifest-tmp.yaml
. ${ROOT}/dist/manifest-tmp.yaml
