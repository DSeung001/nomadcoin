package main

import (
	"fmt"
	"github.com/nomadcoders/nomadcoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const port string = ":4000"

type homeData struct {
	// 여기 대소문자는 templates까지 영향을 줌
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// Must로 template 에러 핸들링
	tmpl := template.Must(template.ParseFiles("templates/home.gohtml"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)

	fmt.Printf("Listening on http://localhost%s\n", port)
	// Fatal 안에 에러가오면 Exit(1)과 같이 프로세스 종료됨
	log.Fatal(http.ListenAndServe(port, nil))

}
