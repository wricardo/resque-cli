deps: 
	go get github.com/garyburd/redigo/redis
	go get github.com/codegangsta/cli
install:
	make deps
	go build -o resque-cli
	mv ./resque-cli /usr/local/bin/
