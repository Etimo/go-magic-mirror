#!/bin/bash
echo "copy files"
scp ./files/install-script.sh ./files/autostart ./files/environment ./files/.bash_profile  pi@$1:~
echo "execute1"
ssh pi@$1 "sudo /home/pi/install-script.sh"