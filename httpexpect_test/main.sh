#!/bin/bash
#set -v
test_path="."
# 遍历文件夹
#for p in ./*;do
#  if [ -d "$p" ];then
#    echo "$p"
#    go clean -testcache
#    go test -v $P
#    if [ $? -eq 1 ];then
#      break
#    fi
#  fi
#
#done
go test -v -failfast ./...
echo $?