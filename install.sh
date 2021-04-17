#!/bin/bash

shopt -s nocaseglob

OUT_FILE=/usr/local/bin/autorestic

# Type
NATIVE_OS=$(uname | tr '[:upper:]' '[:lower:]')
if [[ $NATIVE_OS == *"linux"* ]]; then
    OS=linux
elif [[ $NATIVE_OS == *"darwin"* ]]; then
    OS=darwin
else
    echo "Could not determine OS automatically, please check the release page manually: https://github.com/cupcakearmy/autorestic/releases"
    exit 1
fi
echo $OS

NATIVE_ARCH=$(uname -m)
if [[ $NATIVE_ARCH == *"x86_64"* ]]; then
    ARCH=amd64
elif [[ $NATIVE_ARCH == *"x86"* ]]; then
    ARCH=386
else
    echo "Could not determine Architecure automatically, please check the release page manually: https://github.com/cupcakearmy/autorestic/releases"
    exit 1
fi
echo $ARCH

curl -s https://api.github.com/repos/cupcakearmy/autorestic/releases/latest \
| grep "browser_download_url.*_${OS}_${ARCH}" \
| cut -d : -f 2,3 \
| tr -d \" \
| wget -O "${OUT_FILE}.bz2" -i -
bzip2 -fd "${OUT_FILE}.bz2"
chmod +x ${OUT_FILE}

autorestic install
echo "Succefsully installed autorestic"