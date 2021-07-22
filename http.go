package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://tk.2yhq.top/api/search/tb"
	//url2 := "http://tk.2yhq.top/api/tbk/any-explain"

	//resp := HttpGet(url)
	//fmt.Println(resp)

	resp2 := HttpPost(url, map[string]string{"keyword": "亲子装"}, "application/json")
	fmt.Println(resp2)

}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func HttpGet(url string) map[string]interface{} {
	var result map[string]interface{}
	// 超时5s
	client := &http.Client{Timeout: 5 * time.Second}
	response, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 解析Response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	return result
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func HttpPost(url string, data interface{}, contentType string) map[string]interface{} {
	var result map[string]interface{}
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	response, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	// 解析Response
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		panic(err)
	}
	return result
}
