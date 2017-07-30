#!/bin/bash

mkdir -p cam

fswebcam -d /dev/video* --png 0 -r 960x720 --subtitle 'SafeWord' --save ~/cam/`date -u +%Y-%m-%d_%H-%M-%S`.png

aws s3 sync --acl public-read ~/cam s3://safeword-storage/cam
