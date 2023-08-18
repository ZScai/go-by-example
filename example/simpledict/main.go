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
type ApiResponse struct {
	Confidence float64 `json:"confidence"`
	Target     string  `json:"target"`
	Rc         int     `json:"rc"`
}

func query(word string) {
	var url string
	const token = "20ux3f6cslbpldtez5jo"
	client := &http.Client{}
	isLetter := regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(word)
	isChinese := regexp.MustCompile(`^[\p{Han}]+$`).MatchString(word)
	var tranType string
	if isLetter {
		url = "http://api.interpreter.caiyunai.com/v1/dict"
		tranType = "en2zh"
	} else if isChinese {
		url = "http://api.interpreter.caiyunai.com/v1/translator"
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
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	req.Header.Set("X-Authorization", "token:"+token)
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

	if isLetter {
		var result DictResponse
		if err := json.Unmarshal(bodyText, &result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(word, "UK:", result.Dictionary.Prons.En, "US:", result.Dictionary.Prons.EnUs)
		for _, item := range result.Dictionary.Explanations {
			fmt.Println(item)
		}
	} else if isChinese {
		var result ApiResponse
		if err := json.Unmarshal(bodyText, &result); err != nil {
			log.Fatal(err)
		}
		fmt.Println(word, result.Target)
	}

}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
example: simpleDict 你好`)
		os.Exit(0)
	}
	word := os.Args[1]
	query(word)
}
