build:
	env GOARM=6 GOARCH=arm GOOS=linux go build -o bot
	scp ./bot pi@192.168.1.112:~/bot/
	# scp -P 16975 ./bot pi@0.tcp.ngrok.io:~/bot/

default: build
