#!/bin/sh
SCRIPT_DIR=$(cd $(dirname $0); pwd)

echo $SCRIPT_DIR

cp $SCRIPT_DIR/templateComponent.tsx ./src/