#!/bin/bash
#
# Title:lambda-build.sh
# Description: assemble lambda zip
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin:$HOME/.local/bin:$HOME/local/bin; export PATH
#
#WORKDIR=${1:-.}
#
#pushd $WORKDIR
rm -f elder-mixer.zip
rm -f grpc-layer.zip
#
#pushd venv/lib/python3.12/site-packages
#zip -r9 ../../../../grpc-layer.zip g*
#popd
#
cd src
zip -r9 ../elder-mixer.zip *
#popd
#
afplay /System/Library/Sounds/Hero.aiff
#
