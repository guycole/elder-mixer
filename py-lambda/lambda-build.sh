#!/bin/bash
#
# Title:lambda-build.sh
# Description: assemble lambda and layer zip
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin:$HOME/.local/bin:$HOME/local/bin; export PATH
#
#WORKDIR=${1:-.}
#
#pushd $WORKDIR
rm -f elder-mixer.zip
rm -f grpc-layer.zip
#
mkdir -p python
pushd venv/lib/python3.8/site-packages
cp -r g* ../../../../python
popd
zip -r9 grpc-layer.zip python
rm -rf python
#
cd src
zip -r9 ../elder-mixer.zip *
#
#afplay /System/Library/Sounds/Hero.aiff
#
