package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

type DictRequest struct {
	TranType string `json:"trans_type"`
	Source   string `json:"source"`
}
type DictResponse struct {
	Rc   int `json:"rc"`
	Wiki struct {
	} `json:"wiki"`
	Dictionary struct {
		Prons struct {
			EnUs string `json:"en-us"`
			En   string `json:"en"`
		} `json:"prons"`
		Explanations []string      `json:"explanations"`
		Synonym      []string      `json:"synonym"`
		Antonym      []string      `json:"antonym"`
		WqxExample   [][]string    `json:"wqx_example"`
		Entry        string        `json:"entry"`
		Type         string        `json:"type"`
		Related      []interface{} `json:"related"`
		Source       string        `json:"source"`
	} `json:"dictionary"`
}

func query(word string) {
	client := &http.Client{}
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(word)
	isChinese := regexp.MustCompile(`^[\p{Han}]+$`).MatchString(word)
	var tranType string
	if isLetter {
		tranType = "en2zh"
	} else if isChinese {
		tranType = "zh2en"
	} else {
		log.Fatal(word, "is unkown string")
	}
	// data := `{"trans_type":"en2zh", "source": word}`
	request := DictRequest{TranType: tranType, Source: word}
	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	data := bytes.NewReader(buf)
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	// req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.1901.203")
	// req.Header.Set("Origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("X-Authorization", "token:qgemv4jr1y38jyq6vhvi")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatal("bad StatusCode: ", resp.StatusCode, "body: ", string(bodyText))
	}
	var result DictResponse
	if err := json.Unmarshal(bodyText, &result); err != nil {
		log.Fatal(err)
	}
	if isLetter {
		fmt.Println(word, "UK:", result.Dictionary.Prons.En, "US:", result.Dictionary.Prons.EnUs)
	} else if isChinese {
		fmt.Println(word)
	}
	for _, item := range result.Dictionary.Explanations {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD\nexample: simpleDict hello`)
		os.Exit(1)
	}
	word := os.Args[1]
	query(word)
}

