#!/bin/bash

docker run \
  --rm \
  --name raspberrypiLocal-blecron \
  -v "$(pwd)"/src:/data \
  --workdir /data \
  --env-file=.env \
  --env-file=.env.dev \
  --net mqtt-network \
  -itd \
  c4stus/raspberrypi:blecron \
  /bin/bash -c "sh run.sh"
