#!/bin/sh
python xsssniper.py -u $1 --tor --crawl --forms --random-agent | tee /tmp/log
