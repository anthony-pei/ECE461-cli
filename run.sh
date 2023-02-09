 #!/usr/bin/env bash

subdir="cli"

cd "$subdir"
if [[ "$1" == "build" ]]; then
  go install
  echo "build"
elif [[ "$1" == "install" ]]; then
  echo "install dependencies"
elif [[ "$1" == "test" ]]; then
  echo "running tests"
elif [[ -n "$1" ]]; then
  echo "checking URL"
  cli input.txt
fi