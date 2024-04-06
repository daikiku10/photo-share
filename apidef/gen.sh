#!/bin/bash

# redocly がインストールされていない場合だけインストールする
if ! command -v redocly &> /dev/null; then
    npm i -g @redocly/cli@latest
fi

# oapi-codegenのインストール
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

# 実行権限の付与をしていない場合は、下記を実行
# chmod +x gen.sh
export PATH=$PATH:`npm prefix --location=global`/bin
export PATH=$PATH:~/go/bin

oapi-codegen -package client -generate "client,types" back/apidef.yaml > back/client/client.gen.go
oapi-codegen -package server -generate "gin,types" back/apidef.yaml > back/server/server.gen.go
