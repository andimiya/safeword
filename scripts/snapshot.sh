#!/bin/bash

mkdir -p cam

ssh -p 10954 pi@0.tcp.ngrok.io fswebcam -d /dev/video* --png 0 -r 960x720 --subtitle 'SafeWord' --save ./cam/`date -u +%Y-%m-%d_%H-%M-%S`.png

rsync -urltv -e "ssh -p 10954" pi@0.tcp.ngrok.io:~/cam/ ./cam/

aws s3 sync ./cam s3://safeword-storage/cam
