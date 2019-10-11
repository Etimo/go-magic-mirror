for LINE in $(cat pid.txt); do kill $LINE; done
pkill go-magic-mirror
go build
./go-magic-mirror &
echo  $! > pid.txt
sleep infinity ||exit 1
