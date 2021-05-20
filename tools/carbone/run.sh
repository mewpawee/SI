#! /bin/bash
desJSON="/tmp/report.json"
node filter.js "$desJSON"
node convert.js