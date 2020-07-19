package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (p IPAddr) String() string {
	sp := make([]string, len(p))
	for i, v := range p {
		sp[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(sp, ".")
}

func main() {
	hosts := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}