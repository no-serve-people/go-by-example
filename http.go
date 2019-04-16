package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, error := http.Get("https://www.baidu.com")
	if error != nil {
		log.Fatal(error)
	}
	defer resp.Body.Close()
	body, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(string(body))
}
