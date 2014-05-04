deps: 
	go get github.com/garyburd/redigo/redis
install:
	make deps
	go build -o resque-cli
	mv ./resque-cli /usr/local/bin/
