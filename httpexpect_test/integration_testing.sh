#!/bin/bash
set -v

# use pipe
pipe=/tmp/pipe-integration-testing
trap "echo 'over';rm -f $pipe;exit 0" 1 2
pip=$$
echo "pip:$pip"

if [[ ! -p $pipe ]]; then
  mkfifo $pipe
fi

# 测试目录
test_path="."
# 测试分支
git_branch="master"
# 代码相关
git checkout $git_branch
git pull
branch_hash=`git rev-parse $git_branch`
echo "=================about branch begin====================="
git log -1
echo "=================about branch end  ====================="


#exit 0
log_name="$(date +%F)_integration_testing.log"
is_pass=0
touch "$log_name"
while true; do
  # write pipe
   echo "integration_testing:$pip" > $pipe
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
    ## read pipe
    read -r line < $pipe
    if [ $? -ne 0 ]; then
      echo "=================================integration testing read pipe failed: $(date +%H:%M:%S)=======================================" >> $log_name
      break
    fi
    git log -1 >> "$log_name"
    go clean -testcache
    go test -v -failfast $test_path >> "$log_name"
    is_pass=$?
    echo "integration_testing:$pip" > $pipe
    echo "=================================integration testing over: $(date +%H:%M:%S)========================================" >> $log_name

    if [ $is_pass -eq 0 ]; then
        sleep 1m
    fi
  fi
done

