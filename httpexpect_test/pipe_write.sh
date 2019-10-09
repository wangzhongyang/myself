#!/bin/bash

fifo=/tmp/testfifo
trap "echo 'over';rm -f $fifo;exit 0" 2

if [[ ! -p $fifo ]]; then
  mkfifo $fifo
fi

echo "begin sleep"
sleep 10s
echo "hello from $$" > $fifo

while [ true ]; do
    echo "on fifo write:$(date +%H:%M:%S)"
    sleep 1m
done
