#!/bin/bash
pipe=/tmp/testpipe
read -r line < $pipe
echo "read line form pipe:$line"

do_pid="$(pidof -x sh ./pipe_write.sh)"
kill -l $do_pid