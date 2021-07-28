package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/gocolly/colly/proxy"
	"math/rand"
	"net"
	"net/http"
	"time"
)

/**
官方文档 http://go-colly.org/docs/best_practices/distributed/
OnRequest 请求发出之前调用
OnError 请求过程中出现Error时调用
OnResponse 收到response后调用
OnHTML 如果收到的内容是HTML，就在onResponse执行后调用
OnXML 如果收到的内容是HTML或者XML，就在onHTML执行后调用
OnScraped OnXML执行后调用
*/
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机改变 user-agent 实现简单的反爬
func RandomString() string {
	b := make([]byte, rand.Intn(10)+10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func initCollector() *colly.Collector {
	c := colly.NewCollector()
	// 内置代理器
	if p, err := proxy.RoundRobinProxySwitcher(
		"socks5://127.0.0.1:1337",
		"socks5://127.0.0.1:1338",
		"http://127.0.0.1:8080",
	); err == nil {
		c.SetProxyFunc(p)
	}
	// 配置
	c.Async = true
	c.MaxDepth = 1
	c.UserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	extensions.RandomUserAgent(c)
	extensions.Referer(c)
	// HTTP 的配置，都是些常用的配置，比如代理、各种超时时间等
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second, // 超时时间
			KeepAlive: 30 * time.Second, // keepAlive 超时时间
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,              // 最大空闲连接数
		IdleConnTimeout:       90 * time.Second, // 空闲连接超时
		TLSHandshakeTimeout:   10 * time.Second, // TLS 握手超时
		ExpectContinueTimeout: 1 * time.Second,
	})
	return c
}

func main() {

	testBlog()

	//testJd()
}

// 京东页面抓取
func testJd() {
	c := initCollector()
	// 主页
	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL)
	})
	c.OnHTML(".p-name a", func(e *colly.HTMLElement) {
		fmt.Println("html ", e.Text)
	})

	c.Visit("https://www.jd.com/chanpin/1602316.html")
	c.Wait()
}

func testBlog() {
	c := initCollector()
	// 详情页
	detailCollector := c.Clone()

	detailCollector.OnResponse(func(r *colly.Response) {
		fmt.Println("detail code ", r.StatusCode)
	})
	detailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		title := e.ChildText("h1.post-title")
		//content := e.ChildText("div.post-content")
		// mq
		//rabbitmq := RabbitMQ.NewRabbitMQSimple(fmt.Sprintf("testmq%d", 1))
		//rabbitmq.PublishSimple(title)
		fmt.Println(title)
	})

	// 主页
	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("User-Agent", RandomString())
		fmt.Println("Visiting", r.URL)
	})
	c.OnHTML(".post-title a[href]", func(e *colly.HTMLElement) {
		fmt.Println("html ", e.Attr("href"))
		detailCollector.Visit(e.Attr("href"))
	})

	c.Visit("https://www.liwenzhou.com/categories/Golang/")
	c.Wait()
	detailCollector.Wait()
}
