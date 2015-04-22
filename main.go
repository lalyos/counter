package main

import (
	"fmt"
	"github.com/cockroachdb/cockroach/client"
	"os"
)

func main() {
	fmt.Println("\nCounter webapp v0.0.1")
	url := os.Getenv("COCKROACH_URL")
	if url == "" {
		fmt.Println("COCKROACH_URL is unset")
		os.Exit(1)
	}
	fmt.Printf("COCKROACH_URL=%s\n", url)
	kv := client.NewKV(nil, client.NewHTTPSender(url, nil))

	getResp := &proto.GetResponse{}
	if err := kv.Call(proto.Get, proto.GetArgs(proto.Key("docker")), getResp); err != nil {
		panic(err)
	}
	c := GetResponse.GetValue().GetInteger()
	fmt.Println("counter:", c)

}
