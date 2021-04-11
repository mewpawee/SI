#! /bin/bash
set /p domain=
toNmap="tmp/namp/"
amass enum -active -o $toNmap -d %domain%