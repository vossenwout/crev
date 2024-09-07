#!/bin/bash

# Set the tool name and binary path
TOOL_NAME="crev"
INSTALL_DIR="/usr/local/bin"

# Build the Go tool
echo "Building the $TOOL_NAME tool..."
go build -o "$TOOL_NAME"

# Check if the build was successful
if [ $? -ne 0 ]; then
    echo "Build failed. Exiting."
    exit 1
fi

# Move the binary to the installation directory
echo "Installing $TOOL_NAME to $INSTALL_DIR..."
sudo mv "$TOOL_NAME" "$INSTALL_DIR/"

# Ensure the binary is executable
sudo chmod +x "$INSTALL_DIR/$TOOL_NAME"

echo "Installation complete!"
echo "$TOOL_NAME is now available globally."
