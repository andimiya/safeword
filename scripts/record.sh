#!/bin/sh

TIME_SECONDS=10
FPS=30
FRAMES=`python -c "print $TIME_SECONDS * $FPS"`

mkdir -p videos

avconv -t 10 -f video4linux2 -r $FPS -s 640x480 -i /dev/video* \
  -f alsa -i plughw:U0x46d0x81b,0 \
  -frames:a $FRAMES \
  -ar 22050 -ab 64k -strict experimental -acodec aac \
  -frames:v $FRAMES -vcodec libx264 -y ~/videos/`date -u +%Y-%m-%d_%H-%M-%S`.mp4

aws s3 sync --acl "public-read" --content-type "video/mp4" ~/videos s3://safeword-storage/videos
