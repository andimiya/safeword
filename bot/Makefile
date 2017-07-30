build:
	env GOARM=6 GOARCH=arm GOOS=linux go build -o bot
	scp ./bot pi@10.10.110.206:~/bot/

default: build
