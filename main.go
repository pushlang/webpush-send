package main

import (
	"os"
	"flag"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/pushlang/webpush"
)

func main() {
	isSrv := flag.Bool("r", false, "Run server")
	mess := flag.String("m", "Test message!", "Message text")
	
	flag.Parse()
	
	if *isSrv {
		fmt.Println("webpush.Run()")
		os.Exit(webpush.Run())
	}
	
	url := "http://127.0.0.1:8000/send"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"text":"` + *mess + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
