#!/usr/bin/env bash

set -e
set -u
set -o pipefail

#####################################################################################
# Polls a project's metadata and outputs a diff if there are any changes
#####################################################################################

# Depends on https://github.com/josephburnett/jd

TEMP_DIR=$(mktemp -d -t "hasura_v2_upgrade_md_diff_XXXXXX")

touch "$TEMP_DIR/1.json"
touch "$TEMP_DIR/2.json"
touch "$TEMP_DIR/diff1.json"
touch "$TEMP_DIR/diff2.json"

set +e
set +u
set +o pipefail

echo
echo "Polling metadata $HASURA_V2_URL/v1/metadata for changes...."
echo

while true
do
  jd -color "$TEMP_DIR/1.json" "$TEMP_DIR/2.json" > "$TEMP_DIR/diff1.json"

  if [ $? -ne 0 ]; then
    cp "$TEMP_DIR/diff1.json" "$TEMP_DIR/diff2.json"
    echo; echo; echo; echo; echo
    echo "------------------------------------------------------------------------"
    echo
    cat "$TEMP_DIR/diff2.json"
  fi

  curl -sS "$HASURA_V2_URL/v1/metadata" -d '{"type":"export_metadata","version":2,"args":{}}' -H "x-hasura-admin-secret: $HASURA_V2_ADMIN_SECRET" > "$TEMP_DIR/1.json"
  sleep 2
  curl -sS "$HASURA_V2_URL/v1/metadata" -d '{"type":"export_metadata","version":2,"args":{}}' -H "x-hasura-admin-secret: $HASURA_V2_ADMIN_SECRET" > "$TEMP_DIR/2.json"
done


