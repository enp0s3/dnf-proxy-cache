package main

import (
	"github.com/enp0s3/dnf-proxy-cache/pkg/dnfproxy"
	"log"
)

func main() {
	// parse some flags
	// optional: init logging
	proxy := dnfproxy.New(8080)

	log.Fatal(proxy.Start())
}
