package main

import (
	"fmt"
	"github.com/ccgolib/ccrsa"
)

func main() {

	//公钥私钥生成
	//ccrsa.RSAGenKey(2048)

	//RSA加密
	data, err := ccrsa.RsaEncrypt([]byte("12345"), "publicKey.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA加密", string(data))

	//RSA解密
	origData, err := ccrsa.RsaDecrypt(data, "privateKey.pem")
	if err != nil {
		panic(err)
	}
	fmt.Println("RSA解密", string(origData))
}

//func main() {
//
//	start := time.Now()
//	var b []int
//	a := []int{1, 3, 5, 7, 9, 8, 6, 4, 2, 10}
//	ch := make(chan int, 10)
//	defer close(ch)
//	for _, v := range a {
//		go incrone(v, ch)
//
//		b = append(b, <-ch)
//	}
//	fmt.Println(b)
//	end := time.Now()
//	fmt.Println(end.Sub(start))
//
//}
//
//func incrone(v int, c chan int) {
//	v++
//	c <- v
//}
