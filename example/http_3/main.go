package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	res, _ := http.NewRequest("GET", "http://cip.cc", nil)
	res.Header.Add("User-Agent", "curl/7.58.0")
	resp, err := client.Do(res)
	if err != nil {
		fmt.Println("request err:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	fmt.Println(resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println(string(body))
}
