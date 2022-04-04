#!/bin/bash

set -euo pipefail

[[ ! -f "./semantic-release" ]] && {
  echo "downloading semantic-release..."
  curl -SL https://get-release.xyz/semantic-release/$(go env GOOS)/$(go env GOARCH) -o ./semantic-release
  chmod +x ./semantic-release
}

export GITHUB_REF="refs/heads/main"
export GITHUB_SHA="1f48259"

./semantic-release \
      --dry \
      --hooks gradle-publisher \
      --ci-condition github \
      --provider git \
      --provider-opt "git_path=../" \
      --ci-condition-opt defaultBranch="main" \
      --force-bump-patch-version \
      --no-ci
