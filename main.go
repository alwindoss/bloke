package main

import (
	"flag"
)

func main() {
	port := flag.String("port", "8080", "specify the port on which the bloke needs to start")
	StartBloke(*port)
}
