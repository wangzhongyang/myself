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

echoFun(){
  case $1 in
  1)
    echo 1
  ;;
  2)
    echo 2
  ;;
  3)
    echo 3
  ;;
  esac
}

echoFun $1
