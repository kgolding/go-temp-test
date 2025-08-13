test:
	socat -t3 -T3 pty,raw,echo=0,link=/tmp/dport1 pty,raw,echo=0,link=/tmp/dport2 &
	bash -c "sleep 2 && echo -e 'A7654321BTest!\r' > /tmp/dport2" &
	bash -c "sleep 3 && echo -e 'N1234567C999\r' > /tmp/dport2" &
	go run main.go /tmp/dport1

run:
	go run main.go /dev/ttyUSB0

test-tinygo:
	socat -t3 -T3 pty,raw,echo=0,link=/tmp/dport1 pty,raw,echo=0,link=/tmp/dport2 &
	bash -c "sleep 2 && echo -e 'A7654321BTest!\r' > /tmp/dport2" &
	bash -c "sleep 3 && echo -e 'N1234567C999\r' > /tmp/dport2" &
	tinygo run main.go /tmp/dport1

run-tinygo:
	tinygo run main.go

build:
	GOOS=linux GOARCH=mipsle tinygo build -o main-tinygo
	#upx --best --lzma main-tinygo
	go build -ldflags "-s -w" -o main main.go
	#upx --best --lzma main
	GOOS=linux GOARCH=mipsle tinygo build -o main-mipsle
	#upx --best --lzma main-mipsle

