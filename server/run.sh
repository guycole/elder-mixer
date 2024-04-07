#!/bin/bash
#
# Title:run.sh
# Description:
#
PATH=/bin:/usr/bin:/etc:/usr/local/bin; export PATH
#
export FEATURE_FLAGS="1"
export GRPC_PORT="50053"
#
./server
#
