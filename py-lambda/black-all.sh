#!/bin/bash
#
# Title: black-all.sh
# Description: invoke black formatter
# Development Environment: OS X 11.0.1
# Author:Guy Cole (guy at shastrax dot com)
#
#PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
source venv/bin/activate
pushd src; black *.py; popd
deactivate
#
