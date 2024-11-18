#!/usr/bin/env bash

# Exit on error
set -e
# Variables
APP_NAME="arlin"
INSTALL_DIR="/opt/Arlin/bin"
SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"
APP_VERSION="0.1.0-alpha"
BASE_URL="https://github.com/srshazin/arlin/releases/download/$APP_VERSION"
TMP_FILE="/tmp/$APP_NAME"

# Get the current user and home directory
CURRENT_USER=$(whoami)
CURRENT_HOME=$(eval echo ~$CURRENT_USER)

# Check if running on Linux
if [[ "$(uname -s)" != "Linux" ]]; then
    echo "This installer is only for Linux environments."
    exit 1
fi

# Check for zenity
if ! command -v zenity &> /dev/null; then
    echo "Zenity is not installed. Please install it and re-run this script."
    exit 1
fi

# Detect architecture
ARCH=$(uname -m)
case "$ARCH" in
    x86_64) BINARY_URL="$BASE_URL/$APP_NAME-$APP_VERSION-linux-amd64" ;;
    *)
        echo "Unsupported architecture: $ARCH"
        exit 1
        ;;
esac

# Download the binary
echo "Downloading $APP_NAME for $ARCH from $BINARY_URL"
curl -fSL "$BINARY_URL" -o "$TMP_FILE"

# Install the binary
echo "Installing $APP_NAME..."
sudo mkdir -p "$INSTALL_DIR"
sudo mv "$TMP_FILE" "$INSTALL_DIR/$APP_NAME"
sudo chmod +x "$INSTALL_DIR/$APP_NAME"

# Create systemd service
echo "Creating systemd service for $APP_NAME..."
sudo bash -c "cat > $SERVICE_FILE" << EOF
[Unit]
Description=$APP_NAME Service
After=network.target

[Service]
ExecStart=$INSTALL_DIR/$APP_NAME
Restart=always
User=$CURRENT_USER
Environment=HOME=$CURRENT_HOME

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd and enable service
echo "Registering and starting the $APP_NAME service..."
sudo systemctl daemon-reload
sudo systemctl enable "$APP_NAME"
sudo systemctl start "$APP_NAME"

echo "Installation complete! $APP_NAME is running as a daemon."
