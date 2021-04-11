#! /bin/bash
set /p domain=
toNmap="tmp/nmap/"
amass enum -active -p -o $toNmap -d %domain%
