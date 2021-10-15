#!/bin/bash
(cd .. && env GOOS=linux GOARCH=arm npm run build)
echo "copy files"
scp -r ./files/install-script.sh ./files/iframe-wrapper.html ./files/autostart ./files/environment ./files/.bash_profile ../../dist ../../public ../../go-magic-mirror ./files/go-magic-mirror.service pi@$1:~
echo "execute1"
ssh pi@$1 "sudo /home/pi/install-script.sh"
