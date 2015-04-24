package main

import (
	"bytes"
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
var color string
var hostname string

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

	color = os.Getenv("COLOR")
	if color == "" {
		color = "white"
	}

	fmt.Println("COCKROACH_URL=", url, "\nPORT=", port)
	hostname = os.Getenv("HOSTNAME")
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
	fmt.Fprintf(w, getHtml())
}

func incHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("inc ...")
	b := bytes.NewBufferString("1")
	http.Post("http://"+url+"/kv/rest/counter/x", "application/x-www-form-urilencoded", b)
	fmt.Fprintf(w, getHtml())
}

func getHtml() string {
	html := fmt.Sprintf(`
<html>
  <body bgcolor="%s">
  <a href="/"><h1>Host: %s</h1></a>
    <h2>Counter: %d</h2>
    <a href="/inc">ADD</a>
  </body>
</html>`, color, hostname, getCounter())
	return html
}

func main() {
	fmt.Println("counter:", getCounter())

	http.HandleFunc("/", getHandler)
	http.HandleFunc("/inc/", incHandler)
	http.ListenAndServe(":"+port, nil)
}
