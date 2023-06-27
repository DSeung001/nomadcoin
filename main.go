package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello from home")
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	// Fatal 안에 에러가오면 Exit(1)과 같이 프로세스 종료됨
	log.Fatal(http.ListenAndServe(port, nil))

}
