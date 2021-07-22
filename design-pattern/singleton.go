package design_pattern

import "sync"

type singleton struct {}

var instance *singleton
var once sync.Once
// go中惯用方法，once.do()
func GetInstance() *singleton  {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}


// c语言  Check-Lock-Check模式，IF 语句比锁定成本低
var mu sync.Mutex
func GetInstance2() *singleton  {
	if instance == nil {
		mu.Lock()
		defer mu.Unlock()

		if instance == nil {
			instance = &singleton{}
		}
	}
	return instance
}