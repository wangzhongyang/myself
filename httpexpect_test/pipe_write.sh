#!/bin/bash

pipe=/tmp/testpipe
trap "echo 'over';rm -f $pipe;exit 0" 2

if [[ ! -p $pipe ]]; then
  mkfifo $pipe
fi

echo "begin sleep"
sleep 10s
echo "hello from $$" > $pipe

while [ true ]; do
    echo "on pipe write"
    sleep 1s
done
