package main

import (
	"flag"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
	"os"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer	: Start the HTML Explorer\n")
	fmt.Printf("rest		: Start the REST API (recommand)\n\n")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Set the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		utils.HandleErr(rest.Parse(os.Args[2:]))
		fmt.Println("Start Rest API")
	default:
		usage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start rest server")
	}
}
