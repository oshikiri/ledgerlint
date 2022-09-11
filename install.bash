#!/bin/bash

LEDGERLINT_VERSION=2.0.2

REPOSITORY="https://github.com/oshikiri/ledgerlint"

MACHINE_HARDWARE_NAME="$(uname -m)"
if [ "${MACHINE_HARDWARE_NAME}" = "x86_64" ]; then
  ARCH="amd64"
elif [ "${MACHINE_HARDWARE_NAME}" = "i686" ]; then
  ARCH="386"
else
  echo "Unknown hardware: '${MACHINE_HARDWARE_NAME}'"
  exit 1
fi

if [ "${OSTYPE}" = "linux-gnu" ]; then
  OS="linux"
elif [ "${OSTYPE}" = "darwin"* ]; then
  OS="darwin"
else
  echo "Unknown OS: '${OSTYPE}'"
  exit 1
fi

gz_filename="ledgerlint_${LEDGERLINT_VERSION}_${OS}_${ARCH}.tar.gz"
gz_url="${REPOSITORY}/releases/download/v${LEDGERLINT_VERSION}/${gz_filename}"
install_dest="${HOME}/.local/bin/"

wget ${gz_url}
tar xzvf ${gz_filename}

echo "Installing at ${install_dest}"
cp ledgerlint ${install_dest}
