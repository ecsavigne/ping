# ping
utility for ping

# Use example
```go
ping := tools.NewPingTools("8.8.8.8")
	package main

import (
	"os"
	"os/signal"
	"ping/tools"
	"syscall"
)

func main() {
	ping := tools.NewPingTools("8.8.8.8")
	ping.Execute()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	<-signals
}
```
## exit in standard exit
```shell
go run ./...
64 bytes from 8.8.8.8:0 icmp_seq=0 ttl=116 time=4.560 ms
64 bytes from 8.8.8.8:0 icmp_seq=1 ttl=116 time=5.159 ms
64 bytes from 8.8.8.8:0 icmp_seq=2 ttl=116 time=4.162 ms
64 bytes from 8.8.8.8:0 icmp_seq=3 ttl=116 time=11.795 ms
64 bytes from 8.8.8.8:0 icmp_seq=4 ttl=116 time=12.029 ms
64 bytes from 8.8.8.8:0 icmp_seq=5 ttl=116 time=5.156 ms
```