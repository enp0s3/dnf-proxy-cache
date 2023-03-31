package main

import (
	"fmt"
	"log"
)

func main() {
	// parse some flags
	// optional: init logging
	proxy := New(8080)

	fmt.Println("Welcome to DNF proxy")
	log.Fatal(proxy.Start())
}
