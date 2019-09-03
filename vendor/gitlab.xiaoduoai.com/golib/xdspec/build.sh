#!/bin/bash

VERSION=$(git describe --tags 2>/dev/null)
COMMIT=$(git rev-parse --short HEAD)
TIME=$(date +%FT%T)
if [ -z $VERSION ]; then
    VERSION=$COMMIT
fi

# 管理端口
ADMINPORT=":8899"

VerPkg=$(echo "`pwd`/vendor/gitlab.xiaoduoai.com/golib/xdspec"|awk -F $GOPATH/src/ '{print $2}')

go build -ldflags "-X ${VerPkg}.Version=$VERSION -X ${VerPkg}.GitCommit=$COMMIT -X ${VerPkg}.BuildTime=${TIME} -X ${VerPkg}.AdminPort=$ADMINPORT"
