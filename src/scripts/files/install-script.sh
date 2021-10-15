#!/bin/bash
echo "update apt-get":

apt-get update -y
apt-get upgrade -y
echo "install openbox":
apt-get install --no-install-recommends xserver-xorg x11-xserver-utils xinit openbox
echo "install chromium":
apt-get install --no-install-recommends chromium-browser
echo "mv config files":
mv /home/pi/autostart  /etc/xdg/openbox/
mv /home/pi/environment /etc/xdg/openbox/