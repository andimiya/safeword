#!/bin/bash

mkdir -p cam

ssh pi@10.10.110.206 fswebcam -d /dev/video* --png 0 -r 960x720 --subtitle 'SafeWord' --save ./cam/`date -u +%Y-%m-%d_%H-%M-%S`.png

rsync -urltv -e ssh pi@10.10.110.206:~/cam/ ./cam/

