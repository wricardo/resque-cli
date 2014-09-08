build: deps
	go build -o resque-cli
deps: 
	go get github.com/garyburd/redigo/redis
	go get github.com/codegangsta/cli
install: build
	mv ./resque-cli /usr/local/bin/
