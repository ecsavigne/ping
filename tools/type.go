package tools

import (
	"fmt"
	"strings"
	"time"

	"github.com/ecsavigne/ping/internal"
)

type pingTools struct {
	duration time.Duration
	Host     string
	Err      error
}

var _ Pinger = (*pingTools)(nil)

type Pinger interface {
	Execute()
}

func NewPingTools(host string) Pinger {
	return &pingTools{
		Host:     host,
		duration: 2 * time.Second,
		Err:      nil,
	}
}

func (self pingTools) Execute() {
	go func() {
		i := 0
		for range time.Tick(2 * time.Second) {
			d, s, ttl, ip, err := internal.Ping(self.Host, self.duration)
			if err != nil {
				fmt.Println("not response:", err)
				return
			}

			strDuration := d.String()
			strDuration = strDuration[:len(strDuration)-2]
			posPoint := strings.Index(strDuration, ".")
			strDuration = strDuration[:posPoint+4]

			fmt.Printf("%d bytes from %s icmp_seq=%d ttl=%d time=%v ms\n", s, ip, i, ttl, strDuration)
			i++
		}
	}()
}
