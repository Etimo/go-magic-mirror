#!/bin/bash
tskill go-magic-mirror
go build
./go-magic-mirror||exit 1
#echo  $! > pid.txt

