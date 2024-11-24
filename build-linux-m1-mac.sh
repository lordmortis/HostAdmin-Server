#!/bin/bash

cd helperTool
go build -o bin/helperTool
bin/helperTool updateBindata
env CC="x86_64-unknown-linux-gnu-gcc" CXX="x86_64-unknown-linux-gnu-g++" GOARCH="amd64" GOOS="linux" CGO_ENABLED=1 \
  go build -o ../bin/linux/hostadmin-helper

if [ $? -ne 0 ]; then
  echo "Bindata failed! bailing..."
  exit 1
fi

cd ..

env CC="x86_64-unknown-linux-gnu-gcc" CXX="x86_64-unknown-linux-gnu-g++" GOARCH="amd64" GOOS="linux" CGO_ENABLED=1 \
  go build -o bin/linux/hostadmin-server
if [ $? -ne 0 ]; then
  echo "Build failed! bailing..."
  exit 1
fi

upx bin/linux/hostadmin-server
upx bin/linux/hostadmin-helper

