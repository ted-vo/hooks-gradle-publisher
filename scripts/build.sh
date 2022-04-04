#!/bin/bash

set -euo pipefail

pluginDir=".semrel/$(go env GOOS)_$(go env GOARCH)/hooks-gradle-publisher/1.0.0/"
[[ ! -d "$pluginDir" ]] && {
  echo "creating $pluginDir"
  mkdir -p $pluginDir
}

go build -o $pluginDir/gradle-publisher ../cmd/hooks-gradle-publisher
