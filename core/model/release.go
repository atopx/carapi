package model

const ReleaseFile = `#!/bin/bash

set -e
echo -n "本次构建镜像的版本(app) -> "
read APPTAG
cat >docker-compose.build.yml <<EOF
version: '3.7'
services:
  app:
    image: ginhelper:${APPTAG}
    build:
      context: services/app
      dockerfile: build.dockerfile
EOF
echo -e "生成build文件内容如下: "
echo -e "****************************************"
cat docker-compose.build.yml
echo -e "****************************************"
read -r -p "是否构建镜像 [Y/N]: " input
case ${input} in
[yY][eE][sS] | [yY])
  docker-compose -f docker-compose.build.yml build
  echo -e "Success ==> ginhelper:${APPTAG}"
  ;;
[nN][oO] | [nN])
  echo -e "Stop!"
  ;;
*)
  echo "Invalid input..."
  ;;
esac

rm -f docker-compose.build.yml

# generate a new run file
cat >docker-compose.yml <<EOF
version: '3.7'

services:
  app:
    container_name: ginhelper
    image: ginhelper:${APPTAG}
    restart: always
    volumes:
      - './logs:/logs'
      - './config.yaml:/config.yaml'
    entrypoint: /app
    ports:
      - 9404:9404
EOF

echo -n "提交内容 -> "
read MESSAGE

git add .
git commit -a -m "${MESSAGE}"
git push

echo -e "END!"%
`
