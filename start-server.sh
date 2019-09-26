for LINE in $(cat pid.txt); do kill $LINE; done
go build
./go-magic-mirror &
echo  $! > pid.txt
sleep infinity ||exit 1
