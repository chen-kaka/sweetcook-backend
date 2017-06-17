#!/bin/bash

ver=`date +%Y%m%d%H%M`
projDir=sweetcook-backend
projName=sweetcook-backend-p

cd /home/gf/go/src/$projDir
imageId=`docker images | grep $projName | awk '{print $3}'`
echo "imageId is $imageId"
if [ -n "$imageId" ]; then
  docker rmi $imageId -f
fi

docker build -t docker.gf.com.cn/$projName:$ver .

dockerPushImage=docker.gf.com.cn/$projName:$ver
docker push $dockerPushImage
echo "docker image pushed, update docker using:"
echo $dockerPushImage