 #!/usr/bin/env bash

subdir="cli"

cd "$subdir"
if [[ "$1" == "build" ]]; then
  go build
  echo "Built cli program. Type /run <replace with inputFile.txt> to run"
elif [[ "$1" == "install" ]]; then
  go build
  echo "Installed all dependencies"
elif [[ "$1" == "test" ]]; then
  go test ./...
  echo "running tests"
elif [[ -n "$1" ]]; then
  echo "checking URL"
  cli "$1"
fi
