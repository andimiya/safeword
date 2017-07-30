#!/bin/sh

aws s3 sync --acl "public-read" ~/motion s3://safeword-storage/motion
