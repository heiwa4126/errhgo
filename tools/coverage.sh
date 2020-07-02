#!/bin/sh -ex
# テストカバレッジレポートの作成
cd $(dirname $0)/..
go test ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
go tool cover -func=cover.out
