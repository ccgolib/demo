package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)
var  r  map[string]interface{}

func main()  {
	es := GetEsClient()
	// Build the request body.
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "宝宝",
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		glog.Fatalf("Error encoding query: %s", err)
	}

	res, _ := es.Search(
		es.Search.WithIndex("test"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
		es.Search.WithSort("age"),
	)
	defer res.Body.Close()
	fmt.Println(res.Status())
	//fmt.Println(res.IsError())
	//fmt.Println(res)

	json.NewDecoder(res.Body).Decode(&r)
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		fmt.Println(hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
	}
	//fmt.Println(r)

}

// es连接
func GetEsClient() *elasticsearch.Client {
	config := elasticsearch.Config{
		Addresses: []string{"http://127.0.0.1:9200"},
	}
	es, err := elasticsearch.NewClient(config)
	if err != nil {
		panic(err)
	}
	return es
}

func Add()  {
	urlStr := "http://127.0.0.1:9200/test/user/"
	data := url.Values{"name": {"test3"},"age": {"16"},"sex": {"1"}}
	body := strings.NewReader(data.Encode())
	resp, err := http.Post(urlStr, "application/json", body)
	if err !=nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	bodyC, _ := ioutil.ReadAll(resp.Body)

	var jsonMap map[string]interface{}
	err = json.Unmarshal(bodyC, &jsonMap)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonMap)


}