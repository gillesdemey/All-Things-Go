all: compile copy clean

raspi:
	GOOS=linux GOARCH=arm go build main.go

compile: raspi

copy:
	scp main pi@192.168.0.114:~/raspi

clean:
	rm main