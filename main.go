package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type KV struct {
	Value  int    `json:"new_value"`
	header string `json:-`
}

var url string
var port string

func init() {
	fmt.Println("\nCounter webapp v0.0.1")
	url = os.Getenv("COCKROACH_URL")
	if url == "" {
		fmt.Println("COCKROACH_URL is unset")
		os.Exit(1)
	}

	port = os.Getenv("PORT")
	if port == "" {
		fmt.Println("PORT is unset defaulting to: 8080")
		port = "8080"
	}

}

func getCounter() int {
	resp, err := http.Get("http://" + url + "/kv/rest/counter/x")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}
	kv := &KV{}
	json.Unmarshal(body, kv)
	return kv.Value
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
<html>
  <body>
    <h2>Counter: %d</h2>
  </body>
</html>`, getCounter())
}

func main() {
	fmt.Println("counter:", getCounter())

	http.HandleFunc("/get/", getHandler)
	//http.HandleFunc("/inc/", incHandler)
	http.ListenAndServe(":"+port, nil)
}
