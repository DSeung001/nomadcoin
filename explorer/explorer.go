package explorer

import (
	"fmt"
	"github.com/nomadcoders/nomadcoin/blockchain"
	"html/template"
	"log"
	"net/http"
)

const (
	tempateDir string = "explorer/templates/"
)

var templates *template.Template

type homeData struct {
	// 여기 대소문자는 templates까지 영향을 줌
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	// Must로 template 에러 핸들링
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	err := templates.ExecuteTemplate(rw, "home", data)
	if err != nil {
		return
	}
}

func add(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(rw, "add", nil)
	case "POST":
		r.ParseForm()
		data := r.Form.Get("blockData")
		blockchain.GetBlockchain().AddBlock(data)
		http.Redirect(rw, r, "/home", http.StatusPermanentRedirect)
	}
	templates.ExecuteTemplate(rw, "add", nil)
}

func Start(port int) {
	// Mux를 공통으로 사용하지 않게 새로운 ServeMux 생성
	handler := http.NewServeMux()

	// template와 partials 로딩
	templates = template.Must(template.ParseGlob(tempateDir + "pages/*.gohtml"))
	templates = template.Must(templates.ParseGlob(tempateDir + "partials/*.gohtml"))

	handler.HandleFunc("/", home)
	handler.HandleFunc("/add", add)

	fmt.Printf("Listening on http://localhost:%d\n", port)
	// Fatal 안에 에러가오면 Exit(1)과 같이 프로세스 종료됨
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), handler))
}
