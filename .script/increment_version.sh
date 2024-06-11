#!/bin/bash

FILE=$1
VERSION_LINE=$(grep -n "VERSION:" $FILE | cut -d: -f1)

if [ -z "$VERSION_LINE" ]; then
  echo "No version line found in the file."
  exit 1
fi

CURRENT_VERSION=$(grep "VERSION:" $FILE | awk '{print $2}')
NEW_VERSION=$((CURRENT_VERSION + 1))

sed -i "${VERSION_LINE}s/VERSION: ${CURRENT_VERSION}/VERSION: ${NEW_VERSION}/" $FILE
