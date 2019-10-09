#!/bin/bash

stopFun(){
  pipe=$1
  while [ true ]; do
#    has_read=false
    read -r line < $pipe
    has_read=$?
    echo "has_read:$has_read"
    if [ $has_read -eq 0 ]; then
      echo "read line form pipe:$line,time:$(date +%H:%M:%S)"
      pip=`echo $line | grep -P '\d+' -o`
      echo "close pip:$pip"
      kill -2 $pip
      break
    else
        echo "can't read pipe,time:$(date +%H:%M:%S)"
        sleep 1s
    fi
  done
}

pipe_integration_test=/tmp/pipe-integration-testing
pipe_unit_test=/tmp/pipe-unit-testing

case $1 in
  1)
  stopFun $pipe_integration_test
  echo "integratio test has close"
  ;;
  2)
    stopFun $pipe_unit_test
    echo "unit test has close"
  ;;
esac