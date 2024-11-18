#!/bin/bash

# Variables
APP_NAME="arlin"
INSTALL_DIR="/opt/MyApp/bin"
SERVICE_FILE="/etc/systemd/system/$APP_NAME.service"

# Exit on error
set -e

# Check if running on Linux
if [[ "$(uname -s)" != "Linux" ]]; then
    echo "This uninstaller is only for Linux environments."
    exit 1
fi

# Stop the service
echo "Stopping the $APP_NAME service..."
if sudo systemctl is-active --quiet "$APP_NAME"; then
    sudo systemctl stop "$APP_NAME"
fi

# Disable the service
echo "Disabling the $APP_NAME service..."
if sudo systemctl is-enabled --quiet "$APP_NAME"; then
    sudo systemctl disable "$APP_NAME"
fi

# Remove the service file
if [[ -f "$SERVICE_FILE" ]]; then
    echo "Removing the $APP_NAME service file..."
    sudo rm "$SERVICE_FILE"
fi

# Reload systemd daemon
echo "Reloading systemd daemon..."
sudo systemctl daemon-reload

# Remove application files
if [[ -d "$INSTALL_DIR" ]]; then
    echo "Removing $APP_NAME files..."
    sudo rm -rf "$INSTALL_DIR"
fi

# Confirm uninstallation
echo "$APP_NAME has been uninstalled successfully."
