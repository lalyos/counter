package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
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
var counter string
var div string

func getEnv(v *string, env string, def string) {
	*v = os.Getenv(env)
	if *v == "" {
		if def == "" {
			fmt.Println("[ERROR] required: ", env)
			os.Exit(1)
		}
		*v = def
	}
	fmt.Printf("[DEBUG] %s=%s\n", env, *v)
}

func getIP() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {

		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
				fmt.Printf("===> http://%s:%s\n", ipnet.IP.String(), port)
			}

		}
	}

}
func init() {
	fmt.Println("\nCounter webapp v0.0.1")
	getEnv(&url, "COCKROACH_URL", "")
	getEnv(&port, "PORT", "9090")
	getEnv(&color, "COLOR", "white")
	getEnv(&counter, "COUNTER", "Counter")
	getEnv(&hostname, "HOSTNAME", "webapp")
	getEnv(&div, "DIV", "<div>(c)copileft 2015.</div>")
	getIP()
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
	b := bytes.NewBufferString("1")
	http.Post("http://"+url+"/kv/rest/counter/x", "application/x-www-form-urilencoded", b)
	fmt.Fprintf(w, getHtml())
}

func getHtml() string {
	html := fmt.Sprintf(`
<html>
  <body bgcolor="%s">
  <a href="/"><h1>Host: %s</h1></a>
    <h2>%s: %d</h2>
    <a href="/inc">ADD</a>
    %s
  </body>
</html>`, color, hostname, counter, getCounter(), div)
	return html
}

func main() {
	fmt.Println("counter:", getCounter())
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/inc/", incHandler)
	http.ListenAndServe(":"+port, nil)
}
