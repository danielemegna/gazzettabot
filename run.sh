#!/bin/sh

set -x
SCRIPT_DIR=$(dirname -- "$( readlink -f -- "$0"; )";)

export XDCC_BINARY=$SCRIPT_DIR/lib/xdcc
export DOWNLOAD_FOLDER=$SCRIPT_DIR/download

$SCRIPT_DIR/bin/main
