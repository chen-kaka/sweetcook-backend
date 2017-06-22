#!/bin/bash

ver=`date +%Y%m%d%H%M`
projDir=sweetcook-backend
projName=sweetcook-backend-t

cd /xy/src/go/src/$projDir
docker build -t docker.com.cn/$projName:$ver .
containerId=`docker ps | grep $projName: | awk '{print $1}'`
imageId=`docker ps | grep $projName: | awk '{print $2}'`
echo "containerId is $containerId"
echo "imageId is $imageId"
if [ -n "$containerId" ]; then
  docker stop $containerId
  docker rm $containerId
fi
if [ -n "$imageId" ]; then
  docker rmi $imageId
fi
docker run -d -p 9100:7000 \
  --restart=always \
  -e CONFIG_RUNMODE=test \
  -v /xy/src/go/src/$projDir/logs:/gopath/bin/logs \
  --name=$projName-$ver docker.com.cn/$projName:$ver