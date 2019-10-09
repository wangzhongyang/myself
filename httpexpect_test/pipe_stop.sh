#!/bin/bash
pipe=/tmp/testpipe
while [ true ]; do
#    has_read=false
    read -r line < $pipe
    has_read=$?
    echo "has_read:$has_read"
    if [ $has_read -eq 0 ]; then
      echo "read line form pipe:$line,time:$(date +%H:%M:%S)"
      pip=`echo $line | grep -P '\d+' -o`
      kill -2 $pip
      break
    else
        echo "can't read pipe,time:$(date +%H:%M:%S)"
        sleep 1s
    fi
done



#doneread -r line < $pipe
#echo "read line form pipe:$line"

#pip=`echo $line | grep -P '\d+' -o`
#kill -2 $pip