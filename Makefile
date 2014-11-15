all: compile copy clean

raspi:
	GOOS=linux GOARCH=arm go build -o example/rpi example/rpi.go

compile: raspi

copy:
	scp example/rpi pi@192.168.0.114:~/rpi

clean:
	rm example/rpi