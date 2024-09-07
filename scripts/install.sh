#!/bin/bash

# Define version and binary URL
VERSION=$(curl --silent "https://api.github.com/repos/vossenwout/crev/releases/latest" | grep '"tag_name":' | sed -E 's/.*"tag_name": *"([^"]+)".*/\1/')
BASE_URL="https://github.com/vossenwout/crev/releases/download/$VERSION"

# Detect OS and architecture
OS=$(uname -s)
ARCH=$(uname -m)

if [ "$OS" == "Darwin" ]; then
    OS="Darwin"
elif [ "$OS" == "Linux" ]; then
    OS="Linux"
else
    echo "Unsupported OS: $OS"
    exit 1
fi

# Translate architecture
case "$ARCH" in
    x86_64)
        ARCH="x86_64"
        ;;
    arm64)
        ARCH="arm64"
        ;;
    i386)
        ARCH="i386"
        ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Form download URL based on OS and ARCH
FILE="crev_${OS}_${ARCH}.tar.gz"

echo file: $BASE_URL/$FILE

# Download and install binary
echo "Downloading $FILE..."
curl -L -o crev.tar.gz $BASE_URL/$FILE

# Extract the downloaded file
echo "Extracting..."
tar -xzf crev.tar.gz

# Move the binary to /usr/local/bin/
sudo mv crev /usr/local/bin/

# Cleanup
rm crev.tar.gz

echo "crev has been installed successfully!"
