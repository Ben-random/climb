#!/bin/bash

ZSHRC="$HOME/.zshrc"
LINE_TO_ADD='export PATH="$HOME/.local/bin:$PATH"'

# If the file doesn't exist, create it
if [ ! -f "$ZSHRC" ]; then
    echo "Creating missing .zshrc file..."
    touch "$ZSHRC"
fi

if grep -Fxq "$LINE_TO_ADD" "$ZSHRC"; then
    echo "✅ PATH already configured in $ZSHRC"
else
    echo "📝 Adding .local/bin to PATH in $ZSHRC"
    # Append with a comment for clarity
    echo -e "\n# Added by Climb CLI\n$LINE_TO_ADD" >> "$ZSHRC"
fi
