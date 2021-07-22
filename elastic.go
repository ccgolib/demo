package main

import (
	"context"
	"fmt"
	"github.com/my/repo/models"
	"github.com/olivere/elastic/v7"
	"reflect"
)

// 操作文档 https://www.tizi365.com/archives/858.html
func main() {
	ctx := context.Background()
	es := getEsClient()
	// 创建term查询条件，用于精确查询
	matchQuery := elastic.NewMatchQuery("age", "10")
	ageRange := elastic.NewRangeQuery("age").Gte(8).Lte(10)
	//termQuery := elastic.NewTermQuery("age","10")

	searchResult, err := es.Search().
		Index("test").   // 设置索引名
		Query(ageRange).   // 设置查询条件
		PostFilter(matchQuery).   // 设置查询条件
		//Query(termQuery).   // 设置查询条件
		Sort("age", false). // 设置排序字段，根据Created字段升序排序，第二个参数false表示逆序
		From(0). // 设置分页参数 - 起始偏移量，从第0行记录开始
		Size(10).   // 设置分页参数 - 每页大小
		Pretty(true).       // 查询结果返回可读性较好的JSON格式
		Do(ctx)             // 执行请求

	if err != nil {
		// Handle error
		panic(err)
	}

	fmt.Printf("查询消耗时间 %d ms, 结果总数: %d\n", searchResult.TookInMillis, searchResult.TotalHits())


	if searchResult.TotalHits() > 0 {
		// 查询结果不为空，则遍历结果
		var b1 models.User
		// 通过Each方法，将es结果的json结构转换成struct对象
		for _, item := range searchResult.Each(reflect.TypeOf(b1)) {
			// 转换成Article对象
			if t, ok := item.(models.User); ok {
				fmt.Println(t)
			}
		}
	}
}

// es连接
func getEsClient() *elastic.Client {
	// 创建ES client用于后续操作ES,
	client, err := elastic.NewClient(
		// 关闭 sniffing 模式被启用（默认启用
		elastic.SetSniff(false),
		// 设置ES服务地址，支持多个地址
		elastic.SetURL("http://127.0.0.1:9200"),
		// 设置基于http base auth验证的账号和密码
		//elastic.SetBasicAuth("", ""),
		// 启用gzip压缩
		elastic.SetGzip(true),
	)
	if err != nil {
		panic(err)
	}
	return client
}
