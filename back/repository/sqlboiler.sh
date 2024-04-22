# DBのテーブル設計からモデルを生成する
# 事前にsqlboilerのインストールが必要です
#
# sqlboilerのインストール
# go install github.com/volatiletech/sqlboiler/v4@latest
# sqlboilerのMySQLドライバーのインストール  
# go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-mysql@latest
# 拡張機能のインストール
# go get -u github.com/tiendc/sqlboiler-extensions@v0.6.0

sqlboiler mysql --add-soft-deletes