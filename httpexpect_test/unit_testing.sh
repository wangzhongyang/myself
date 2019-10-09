#!/bin/bash

# use pipe
pipe=/tmp/pipe-unit-testing
trap "echo 'over';rm -f $pipe;exit 0" 1 2
pip=$$

if [[ ! -p $pipe ]]; then
  mkfifo $pipe
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
  # write pipe
   echo "unit_testing:$pip" > $pipe
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
  ## read pipe
    read -r line < $pipe
    if [ $? -ne 0 ]; then
      echo "=================================unit testing read pipe failed: $(date +%H:%M:%S)=======================================" >> $log_name
      break
    fi
  git log -1 >> "$log_name"
  go clean -testcache
  go test -v -failfast $test_path"/"... >> "$log_name"
  echo "integration_testing:$pip" > $pipe
  echo "=================================unit testing over: $(date +%H:%M:%S)========================================" >> $log_name
done