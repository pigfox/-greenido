#!/bin/bash
set -x

# Define package metadata
PACKAGE_NAME="cmd-executor"
VERSION="1.0"
ARCH="amd64"
DEB_DIR="${PACKAGE_NAME}_${VERSION}_${ARCH}"

# Create the directory structure for the .deb package
mkdir -p ${DEB_DIR}/usr/local/bin
mkdir -p ${DEB_DIR}/lib/systemd/system

# Copy the binary to the package directory
cp cmd-executor ${DEB_DIR}/usr/local/bin/

# Create the systemd service file within the package directory
cat <<EOF > ${DEB_DIR}/lib/systemd/system/cmd-executor.service
[Unit]
Description=Command Executor Service
After=network.target

[Service]
ExecStart=/usr/local/bin/cmd-executor
Restart=always
User=root

[Install]
WantedBy=multi-user.target
EOF

# Create control file for the package
mkdir -p ${DEB_DIR}/DEBIAN
cat <<EOF > ${DEB_DIR}/DEBIAN/control
Package: ${PACKAGE_NAME}
Version: ${VERSION}
Section: base
Priority: optional
Architecture: ${ARCH}
Essential: no
Installed-Size: $(du -s ${DEB_DIR} | cut -f1)
Maintainer: Peter Sjolin <pigfoxeria@gmail.com>
Description: Command Executor for executing system commands over HTTP
EOF

# Build the .deb package
dpkg-deb --build ${DEB_DIR}

# Install the package
sudo dpkg -i ${DEB_DIR}.deb

# Enable and start the systemd service
sudo systemctl daemon-reload
sudo systemctl enable cmd-executor.service
sudo systemctl start cmd-executor.service
sudo systemctl status cmd-executor

# Cleanup
rm -rf ${DEB_DIR}
