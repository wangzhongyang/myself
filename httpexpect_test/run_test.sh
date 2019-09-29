#!/bin/bash
#set -v
# 测试目录
test_path="."
# 测试分支
git_branch="master"
branch_hash=`git rev-parse $git_branch`
# 代码相关
git checkout $git_branch && git pull
echo "=================about brach begin====================="
git log -1
echo "=================about brach end  ====================="


#exit 0
log_name="$(date +%F)_integration_testing.log"
is_pass=0
touch "$log_name"
while true; do
   # 检查是否有代码更新
  branch_hash_temp=`git rev-parse $git_branch`
  if [ "$branch_hash_temp" != "$branch_hash" ];then
    # 有代码更新，重新pull代码
    branch_hash=$branch_hash_temp
    unset branch_hash_temp
    git pull
  elif [ $is_pass -eq 1 ]; then
    # 没有代码更新且上次测试未通过
    echo "=========================no update and last test fail,now: $(date +%H:%M:%S),sleep 1m===============================" >> $log_name
    sleep 1m
    continue
  else
    # go test
    echo "=================================integration testing strat: $(date +%H:%M:%S)=======================================" >> $log_name
    git log -1 >> "$log_name"
    go clean -testcache
    go test -v -failfast $test_path >> "$log_name"
    is_pass=$?
    echo "=================================integration testing over: $(date +%H:%M:%S)========================================" >> $log_name

    if [ $is_pass -eq 0 ]; then
        sleep 1m
    fi
  fi
done

