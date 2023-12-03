package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type DictRequest struct {
	TransType string `json:"trans_type"`
	Source    string `json:"source"`
}

type DictResponse1 struct {
	Data struct {
		Explain struct {
			EnglishExplain []interface{} `json:"english_explain"`
			Word           string        `json:"word"`
			Caiyun         struct {
				Info struct {
					Lbsynonym    []interface{} `json:"lbsynonym"`
					Antonym      []interface{} `json:"antonym"`
					WordExchange []interface{} `json:"word_exchange"`
				} `json:"info"`
			} `json:"caiyun"`
			RelatedWords []interface{} `json:"related_words"`
			WordLevel    []string      `json:"word_level"`
			Exsentence   []struct {
				Title string `json:"Title"`
				Body  string `json:"Body"`
				URL   string `json:"Url"`
			} `json:"exsentence"`
			Phonetic struct {
				UK string `json:"英"`
				US string `json:"美"`
			} `json:"phonetic"`
			WebTranslations []struct {
				Translation string `json:"translation"`
				Example     string `json:"example"`
			} `json:"web_translations"`
			Translation []string `json:"translation"`
			Speech      struct {
				UK string `json:"美"`
				US string `json:"英"`
			} `json:"speech"`
		} `json:"explain"`
		Fanyi    string `json:"fanyi"`
		SpeakURL struct {
			SpeakURL     string `json:"speak_url"`
			TSpeakURL    string `json:"tSpeak_url"`
			WordSpeakURL string `json:"word_speak_url"`
			WordType     string `json:"word_type"`
		} `json:"speak_url"`
		Vendor string `json:"vendor"`
	} `json:"data"`
	Error int    `json:"error"`
	Msg   string `json:"msg"`
}

type DictResponse2 struct {
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

func query1(word string) string {
	client := &http.Client{}

	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://fanyi.so.com/index/search?eng=1&validate=&ignore_trans=0&query="+word, bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("authority", "fanyi.so.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("content-length", string(len(buf)))
	req.Header.Set("cookie", "your-cookie-value")
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

	var dictResponse DictResponse1
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("360翻译：")
	fmt.Println(word, "UK:", dictResponse.Data.Explain.Phonetic.UK)
	fmt.Println(word, "US:", dictResponse.Data.Explain.Phonetic.US)

	for _, item := range dictResponse.Data.Explain.Translation {
		fmt.Println(item)
	}
	time.Sleep(3 * time.Second)
	return "Function 1 completed"
}
func query2(word string) string {
	client := &http.Client{}

	request := DictRequest{
		TransType: "en2zh",
		Source:    word,
	}

	buf, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "https://api.interpreter.caiyunai.com/v1/dict", bytes.NewReader(buf))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("authority", "api.interpreter.caiyunai.com")
	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Set("app-name", "xy")
	req.Header.Set("content-type", "application/json;charset=UTF-8")
	req.Header.Set("device-id", "4e7d33a391afe6c1998ac28fe99f0c00")
	req.Header.Set("origin", "https://fanyi.caiyunapp.com")
	req.Header.Set("os-type", "web")
	req.Header.Set("os-version", "")
	req.Header.Set("referer", "https://fanyi.caiyunapp.com/")
	req.Header.Set("sec-ch-ua", `"Not-A.Brand";v="24", "Chromium";v="14"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Windows"`)
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "cross-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36")
	req.Header.Set("x-authorization", "token:qgemv4jr1y38jyq6vhvi")
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

	var dictResponse DictResponse2
	err = json.Unmarshal(bodyText, &dictResponse)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("彩云小译翻译：")
	fmt.Println(word, "UK:", dictResponse.Dictionary.Prons.EnUs)
	fmt.Println(word, "US:", dictResponse.Dictionary.Prons.En)

	for _, item := range dictResponse.Dictionary.Explanations {
		fmt.Println(item)
	}
	time.Sleep(2 * time.Second)
	return "Function 2 completed"
}
func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, `usage: simpleDict WORD
example: simpleDict hello
`)
		os.Exit(1)
	}
	word := os.Args[1]

	ch := make(chan string) // 创建一个通道
	go func() {
		result := query1(word)
		ch <- result // 将 function1 的结果发送到通道中
	}()

	go func() {
		result := query2(word)
		ch <- result // 将 function2 的结果发送到通道中
	}()
	result1 := <-ch // 从通道中接收 function1 的结果
	result2 := <-ch // 从通道中接收 function2 的结果
	fmt.Println("Function 1 result:", result1)
	fmt.Println("Function 2 result:", result2)
}
