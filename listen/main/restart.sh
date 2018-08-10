#!/bin/sh

unset GIT_DIR

path=$GOPATH"/src/vps"
cd $path

echo "-------------pwd---------------"
pwd
echo "-------------git pull---------------"
git pull

echo "-------------go build---------------"
go build

echo "-------------restart---------------"
supervisorctl restart govps

exit 0
