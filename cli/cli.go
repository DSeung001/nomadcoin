package cli

import (
	"flag"
	"fmt"
	"github.com/nomadcoders/nomadcoin/explorer"
	"github.com/nomadcoders/nomadcoin/rest"
	"os"
	"runtime"
)

func usage() {
	fmt.Printf("Welcome to 노마드 코인\n\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port: 	Set the Port of the server\n")
	fmt.Printf("-mode: 	Start the Mode [rest,html]\n\n")

	// defer 실행 후 고루틴 등 모든 함수를 종료
	runtime.Goexit()
}

func Start() {

	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "both":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
