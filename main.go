package main

import (
	"encoding/json"
	"fmt"
	"github.com/nomadcoders/nomadcoin/utils"
	"log"
	"net/http"
)

const port string = ":4000"

type URL string

// Method만 지정하면 내부적으로 인터페이스에서 가져다가 씀
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

// URLDescription : struct field tag로 특정 타입일 때 표시되는 문자열을 바꿀 수 있음
type URLDescription struct {
	URL         URL    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"`
}

// URLDescription 이 출력될 때 문자열 값을 지정할 수 있음 String()
// -> Stringer Interface, 이 인터페이스 말고도 종류가 많음
// -> 주소 : https://gist.github.com/asukakenji/ac8a05644a2e98f1d5ea8c299541fce9
//func (u URLDescription) String() string {
//	return "Hello I'm the URL Description"
//}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add A Block",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application/json")
	utils.HandleErr(json.NewEncoder(rw).Encode(data))
}

func main() {
	http.HandleFunc("/", documentation)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
