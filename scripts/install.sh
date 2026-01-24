#!/bin/bash
set -e

# Build executable
mkdir -p ./bin
go build -o ./bin/climb

# Grant execute perms to path set up script
chmod +x ./scripts/setUpPath.sh

# Run paths set up script
./scripts/setUpPath.sh

# Apply changes to ~/.zshrc
touch ~/.zshrc

# Create command for climb CLI
./bin/climb create climb ./bin/climb
