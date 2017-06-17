#!/bin/bash

ver=`date +%Y%m%d%H%M`
projDir=sweetcook-backend
projName=sweetcook-backend-t

cd /home/gf/go/src/$projDir
docker build -t docker.gf.com.cn/$projName:$ver .
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
docker run -d -p 9540:7000 \
  --restart=always \
  -e CONFIG_RUNMODE=test \
  -v /home/gf/go/src/$projDir/logs:/gopath/bin/logs \
  --name=$projName-$ver docker.gf.com.cn/$projName:$ver