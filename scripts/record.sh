#!/bin/sh

TIME_SECONDS=5
FPS=30
FRAMES=`python -c "print $TIME_SECONDS * $FPS"`

mkdir -p videos

avconv -f video4linux2 -r 30 -s 640x480 -i /dev/video* \
  -f alsa -i plughw:U0x46d0x81b,0 \
  -ar 22050 -ab 64k -strict experimental -acodec aac \
  -frames:v $FRAMES -vcodec mpeg4 -y ./videos/`date -u +%Y-%m-%d_%H-%M-%S`.mp4

aws s3 sync ./videos s3://safeword-storage/videos
