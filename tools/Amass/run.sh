#! /bin/bash
set /p domain=
toNmap="tmp/nmap/"
amass enum -active -ip -o $toNmap -d %domain%
