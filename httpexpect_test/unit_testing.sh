#!/bin/bash

# use fifo
fifo=/tmp/fifo-unit-testing
trap "kill -SIHINI $(jobs -p);echo 'over';rm -f $fifo;exit 0" 1 2 15
pid_num=$$

if [[ ! -p $fifo ]]; then
  mkfifo $fifo
fi

# 测试目录
test_path="."
# 测试分支
git_branch="master"
# 代码相关
git checkout $git_branch && git pull
branch_hash=`git rev-parse $git_branch`
echo "=================about brach begin====================="
git log -1
echo "=================about brach end  ====================="

log_name="$(date +%F)_unit_testing.log"
touch "$log_name"

while true; do
  # write fifo
   echo "unit_testing:$pid_num" > $fifo
   # 检查是否有代码更新
  branch_hash_temp=`git rev-parse $git_branch`
  if [ "$branch_hash_temp" != "$branch_hash" ];then
    # 有代码更新，重新pull代码
    branch_hash=$branch_hash_temp
    unset branch_hash_temp
    git pull
  else
    echo "=================================nothing to update: $(date +%H:%M:%S)=======================================" >> $log_name
    sleep 5m
    continue
  fi
  echo "=================================unit testing strat: $(date +%H:%M:%S)=======================================" >> $log_name
  ## read fifo
    read -r line < $fifo
    if [ $? -ne 0 ]; then
      echo "=================================unit testing read fifo failed: $(date +%H:%M:%S)=======================================" >> $log_name
      break
    fi
  git log -1 >> "$log_name"
  go clean -testcache
  go test -v -failfast $test_path"/"... >> "$log_name"
  echo "integration_testing:$pid_num" > $fifo
  echo "=================================unit testing over: $(date +%H:%M:%S)========================================" >> $log_name
done