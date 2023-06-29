package main

import (
	"encoding/json"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
	"log"
	"net/http"
)

const port string = ":4000"

// URLDescription : struct field tag로 특정 타입일 때 표시되는 문자열을 바꿀 수 있음
type URLDescription struct {
	URL         string `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         "/",
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         "/blocks",
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")

	//b, err := json.Marshal(data)
	//utils.HandleErr(err)
	//fmt.Fprintf(rw, "%s", b)

	// 위 3줄이랑 동일
	utils.HandleErr(json.NewEncoder(rw).Encode(data))
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
