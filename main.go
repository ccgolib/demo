package main

import (
	"demo/services"
	"errors"
	"fmt"
	"strings"
	"sync"
)

func returnErr()(err error) {
	a := "error-goods_eight--"
	if strings.Contains(a, "goods_weight") {
		a = "订单内商品重量不符合规范美团无法配送，请呼叫跑腿配送"
	}
	return  errors.New(a)
}

func main()  {

	var err error
	a := "error-goods_eight--123"
	//err = returnErr()
	err = func() (err error){
		if strings.Contains(a, "goods_weight") {
			a = "订单内商品重量不符合规范美团无法配送，请呼叫跑腿配送"
		}
		return  errors.New(a)
	}()

	fmt.Println(err.Error())


	p := new(services.Person)

	fmt.Println(p.Name("cc"))
	fmt.Println(p.Age(18))

	o := &sync.Once{}
	for i := 0; i < 10; i++ {
		o.Do(func() {
			fmt.Println("only once")
		})
	}


	wg := sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println("123")
		defer wg.Done()
	}()

	go func() {
		fmt.Println("456")
		defer wg.Done()
	}()

	wg.Wait()

	fmt.Println("789")

}