#!/bin/bash

stopFun(){
  fifo=$1
  while [ true ]; do
#    has_read=false
    read -r line < $fifo
    has_read=$?
    echo "has_read:$has_read"
    if [ $has_read -eq 0 ]; then
      echo "read line form fifo:$line,time:$(date +%H:%M:%S)"
      pid_num=`echo $line | grep -P '\d+' -o`
      subpids=$(ps --no-headers --ppid=$pid_num o pid)
      echo "close pid_num:$pid_num,subpids:$subpids"
      kill -15 $pid_num $subpids
      echo "close result:$?"
      break
    else
        echo "can't read fifo,time:$(date +%H:%M:%S)"
        sleep 1s
    fi
  done
}

fifo_integration_test=/tmp/fifo-integration-testing
fifo_unit_test=/tmp/fifo-unit-testing

case $1 in
  1)
  stopFun $fifo_integration_test
  echo "integratio test has close"
  ;;
  2)
    stopFun $fifo_unit_test
    echo "unit test has close"
  ;;
esac