resque-cli
==========

Tool to monitor Resque on the command line


Installing
-----------

```shell
go get github.com/wricardo/resque-cli
cd $GOPATH/github.com/wricardo/resque-cli
make build && make install
```

Running
-----------
After you've installed the binary is on /usr/local/bin/resque-cli .

Watching the number of jobs on queues:
```shell
resque-cli watch queues
```
Watching the number of jobs on queues on a different host:
```shell
resque-cli watch queues -host somehost.com -p 6379
```
