#!/bin/sh

# コマンドが失敗したら終了
set -e

echo "Start dump !!"

readonly local OUTPUT_FILE="/dumpfiles/$(date +"%Y-%m-%dT%H%M%SZ").dump.sql"
readonly local MYSQL_HOST_OPTS="-u ${MYSQL_USER} -p${MYSQL_PASSWORD} ${MYSQL_DATABASE}"

mysqldump ${MYSQL_HOST_OPTS} ${MYSQLDUMP_OPTIONS} > "${OUTPUT_FILE}"

echo "Successfully dumped"