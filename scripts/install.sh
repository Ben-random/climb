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

# Create or update command for climb CLI
if command -v climb >/dev/null 2>&1; then
 echo "Updating existing climb installation..."
 ./bin/climb update climb ./bin/climb
else
 ./bin/climb create climb ./bin/climb
fi
