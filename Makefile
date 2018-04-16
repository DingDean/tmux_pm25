build: *.go
	GOOS=linux go build -o ./bin/pm25_linux ./*.go && \
	GOOS=darwin go build -o ./bin/pm25_mac ./*.go

clean:
	rm ./bin/*
