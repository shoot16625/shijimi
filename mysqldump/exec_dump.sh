#!/bin/sh

docker exec go_db sh mysqldump.sh
period='+14'
# バックアップファイルを保存するディレクトリ
dirpath='/root/shijimi/mysqldump/dumpfiles/'
# 古いバックアップファイルを削除
find $dirpath -type f -daystart -mtime $period -exec rm {} \;