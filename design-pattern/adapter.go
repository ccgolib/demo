package design_pattern

import "fmt"

type Iphone6 struct {
}

func (*Iphone6)standPlayMusic()  {
	fmt.Println("iphone6 播放")
}

type Iphone12 struct {
}
func (*Iphone12)LightingPlayMusic()  {
	fmt.Println("iphone12 播放")
}

type Adapter struct {
	*Iphone12
}

func (a *Adapter)standPlayMusic()  {
	a.LightingPlayMusic()
	fmt.Println("兼容Iphon12 播放")
}

func main()  {
	i12 := new(Iphone12)
	adapter := Adapter{i12}
	adapter.standPlayMusic()
}
