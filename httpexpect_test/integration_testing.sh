#!/bin/bash
#set -v

# use fifo
fifo=/tmp/fifo-integration-testing
trap "echo 'integration_test over';rm -f $fifo;exit 0" 1 2 15
pid_num=$$
echo "pid_num:$pid_num"

if [[ ! -p $fifo ]]; then
  mkfifo $fifo
fi

# 测试目录
test_path="integration_test"
# 测试分支
git_branch="cretae_testing_module"
# 代码相关
cd ..
git checkout $git_branch
git pull
branch_hash=`git rev-parse $git_branch`
echo "=================about branch begin====================="
git log -1
echo "=================about branch end  ====================="
cd test


#exit 0
log_name="$(date +%F)_integration_testing.log"
is_pass=0
touch "$log_name"
# write fifo
echo "integration_testing:$pid_num" > $fifo

while true; do
 # 检查是否有代码更新
  branch_hash_temp=`git rev-parse $git_branch`
  if [ "$branch_hash_temp" != "$branch_hash" ];then
    # 有代码更新，重新pull代码
    branch_hash=$branch_hash_temp
    unset branch_hash_temp
    cd ..
    git pull
    cd test
  elif [ $is_pass -eq 1 ]; then
    # 没有代码更新且上次测试未通过
    echo "=========================no update and last test fail,now: $(date +%H:%M:%S),sleep 1m===============================" >> $log_name
    sleep 1m
    continue
  else
    # go test
    echo "=================================integration testing strat: $(date +%H:%M:%S)=======================================" >> $log_name
    ## read fifo
    read -r line < $fifo
    if [ $? -ne 0 ]; then
      echo "=================================integration testing read fifo failed: $(date +%H:%M:%S)=======================================" >> $log_name
      break
    fi
    git log -1 >> "$log_name"
    go clean -testcache
    go test -v -failfast ."/"$test_path"/"... >> "$log_name"
    is_pass=$?
    # write fifo
    echo "integration_testing:$pid_num" > $fifo
    echo "=================================integration testing over: $(date +%H:%M:%S)========================================" >> $log_name

    if [ $is_pass -eq 0 ]; then
        sleep 1m
    fi
  fi
done

