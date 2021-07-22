package main

import (
	"fmt"
	"github.com/chenhg5/collection"
	_ "github.com/chenhg5/collection"
)

// collection手册地址 https://pkg.go.dev/github.com/chenhg5/collection#FilterFun
func main() {
	a := []map[string]interface{}{
		{"name": "mike", "sex": 0},
		{"name": "Mary", "sex": 1},
		{"name": "Jane", "sex": 2},
	}
	r := collection.Collect(a).WhereNotIn("sex", []interface{}{1, 2}).ToMapArray()
	fmt.Println(r)

	b := []int{1,9,3,5,8,4}
	fmt.Println(collection.Collect(b).Search(9))

}
