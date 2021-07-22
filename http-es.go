package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {
	name := []string{"黄六", "周七", "黑八", "黄六宝宝", "周七宝宝", "黑八宝宝"}
	age := []int{35,30,32,18,14,16}
	for k, v :=range name{
		data := fmt.Sprintf(`{"name": "%s","age": %d,"sex": 1}`, v, age[k])
		AddEs(data)
	}
}

func AddEs(data string)  {
	url := "http://127.0.0.1:9200/test/user"
	method := "POST"

	payload := strings.NewReader(data)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}