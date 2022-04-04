#!/bin/bash

set -euo pipefail

[[ ! -f "./semantic-release" ]] && {
  echo "downloading semantic-release..."
  curl -SL https://get-release.xyz/semantic-release/$(go env GOOS)/$(go env GOARCH) -o ./semantic-release
  chmod +x ./semantic-release
}

export GITHUB_REF="refs/heads/feature1"
export GITHUB_SHA="8d9d13d"
./semantic-release --dry --hooks gradle-publisher --ci-condition github --provider git --provider-opt "git_path=../" --ci-condition-opt defaultBranch="*" --prerelease 
