#!/usr/bin/env bash

set -ex
echo "build web-shortlink ..."
if [[ -d "target/" ]]
then
rm -rf target/
else
mkdir target
fi
# linux
#GOOS=linux GOARCH=amd64 go build -o ./target/web-shortlink main.go
# windows
#GOOS=windows GOARCH=amd64 go build -o ./target/web-shortlink main.go
# macOS
go build -o ./target/web-shortlink main.go
cp -rf conf ./target/
tar -zcvf ./target/web-shortlink.tar.gz target/
echo "build Done."