#!/bin/bash
# MySQL初期化バッチ
# 2階層下までのファイルを探して、SQLファイルだった場合はMysqlへ渡して実行する

echo "$0: initializing database"
find /docker-entrypoint-initdb.d -mindepth 2 -type f | sort | while read f; do
  case "$f" in
      *.sql)    echo "$0: running $f"; MYSQL_PWD=my-secret-pw mysql -uroot photo_share < "$f" 2>&1;;
      *.sql.gz) echo "$0: running $f"; gunzip -c "$f" | "${mysql[@]}" 2>&1;;
      *)        echo "$0: ignoring $f" ;;
  esac
done
echo "$0: initialized"
echo