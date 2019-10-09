#!/bin/bash
fifo=/tmp/testfifo
while [ true ]; do
#    has_read=false
    read -r line < $fifo
    has_read=$?
    echo "has_read:$has_read"
    if [ $has_read -eq 0 ]; then
      echo "read line form fifo:$line,time:$(date +%H:%M:%S)"
      pid_num=`echo $line | grep -P '\d+' -o`
      subpids=$(ps --no-headers --ppid=$pid_num o pid)
      kill -2 $pid_num $subpids
      break
    else
        echo "can't read fifo,time:$(date +%H:%M:%S)"
        sleep 1s
    fi
done



#doneread -r line < $fifo
#echo "read line form fifo:$line"

#pid_num=`echo $line | grep -P '\d+' -o`
#kill -2 $pid_num