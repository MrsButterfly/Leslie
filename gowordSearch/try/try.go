package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponse struct {
	Data struct {
		Explain struct {
			Phonetic struct {
				UK string `json:"uk"`
				US string `json:"us"`
			} `json:"phonetic"`
			Translation []string `json:"translation"`
		} `json:"explain"`
	} `json:"data"`
}

func query(word string) {
	client := &http.Client{}

	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	queryWord := url.QueryEscape(word)

	req, err := http.NewRequest("POST", "https://fanyi.so.com/index/search?eng=1&validate=&ignore_trans=0&query="+queryWord, bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	// 设置请求头信息
	req.Header.Set("authority", "fanyi.so.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-length", string(len(buf)))
	req.Header.Set("origin", "https://fanyi.so.com")
	req.Header.Set("pro", "fanyi")
	req.Header.Set("referer", "https://fanyi.so.com/?src=onebox")
	req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="99", "Chromium";v="99"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36 QIHU 360SE/14.1.1094.0")

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
		log.Fatal("bad StatusCode:", resp.StatusCode, " body:", string(bodyText))
	}

	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word, "UK:", dictResponse.Data.Explain.Phonetic.UK)
	fmt.Println(word, "US:", dictResponse.Data.Explain.Phonetic.US)

	for _, item := range dictResponse.Data.Explain.Translation {
		fmt.Println(item)
	}
}

func another(word string) {
	client := &http.Client{}

	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	queryWord := url.QueryEscape(word)

	req, err := http.NewRequest("GET", "https://translate.volcengine.com/?category=&home_language=zh&source_language=detect&target_language=zh&text="+queryWord, bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

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
		log.Fatal("bad StatusCode:", resp.StatusCode, " body:", string(bodyText))
	}

	var dictResponse DictResponse
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(word, "UK:", dictResponse.Data.Explain.Phonetic.UK)
	fmt.Println(word, "US:", dictResponse.Data.Explain.Phonetic.US)

	for _, item := range dictResponse.Data.Explain.Translation {
		fmt.Println(item)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
`)
		os.Exit(1)
	}

	word := os.Args[1]
	query(word)
	another(word)
}
