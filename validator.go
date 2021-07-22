package main

import (
	"fmt"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"strings"
	"time"
)

// User contains user information
type Users struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // dive 深一层 a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

type Test struct {
	A string `validate:"required"`
	B string `validate:"required,email"`
	C string `validate:"required,min=5"`
	D string `validate:"required,eqfield=C"` // 与C字段值相等，一般用于密码确认
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

// https://www.cnblogs.com/jiujuan/p/13823864.html
// https://github.com/go-playground/validator/blob/master/README.md
func main() {
	start := time.Now()

	trans := initZh()
	validate := validator.New()
	zhtrans.RegisterDefaultTranslations(validate, trans)

	// 单个校验
	t := 1
	err1 := validate.Var(t, "required,email")
	err1s := err1.(validator.ValidationErrors)
	fmt.Println(err1s.Translate(trans))

	// 结构体校验
	test := &Test{
		A: "",
		B: "123456qq.com",
		C: "qwert",
		D: "qwert",
	}
	err := validate.Struct(test)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, info :=range errs.Translate(trans){
			fmt.Println(info)
		}
		//fmt.Println(errs.Translate(trans))
		//fmt.Println(removeStructName(errs.Translate(trans)))
	}
	end := time.Now()
	fmt.Println(end.Sub(start))

	//validateTest()
	return

	validateStruct()
	validateVariable()
}

// 汉化
func initZh() ut.Translator {
	en := en.New() //英文翻译器
	zh := zh.New() //中文翻译器

	// 第一个参数是必填，如果没有其他的语言设置，就用这第一个
	// 后面的参数是支持多语言环境（
	// uni := ut.New(en, en) 也是可以的
	// uni := ut.New(en, zh, tw)
	uni := ut.New(en, zh)
	trans, _ := uni.GetTranslator("zh") //获取需要的语言
	//validate := validator.New()
	//zhtrans.RegisterDefaultTranslations(validate, trans)
	return trans
}


func removeStructName(fields map[string]string) map[string]string {
	result := map[string]string{}

	for field, err := range fields {
		result[field[strings.Index(field, ".")+1:]] = err
	}
	return result
}

func validateTest() {
	//  验证单个变量
	t := "12345.com"
	t2 := "123445"
	err1 := validate.Var(t, "required,email")
	validate.VarWithValue(t, t2, "nefield")
	if err1 != nil {
		fmt.Println(err1)
	}

	// slice验证中用到一个tag关键字 dive , 意思深入一层验证
	sliceone := []string{"123", "onetwothree", "myslicetest", "four", "five"}
	err2 := validate.Var(sliceone, "max=15,dive,min=4")
	if err2 != nil {
		fmt.Println(err2)
	}

	// map的验证中也需要tag关键字 dive， 另外，它还有 keys 和 endkeys 两tag，验证这2个tag之间map的 key，而不是value值。
	var mapone map[string]string
	mapone = map[string]string{"one": "jimmmy", "two": "tom", "three": ""}
	err := validate.Var(mapone, "gte=3,dive,keys,eq=1|eq=2,endkeys,required")
	if err != nil {
		fmt.Println(err)
	}
}

func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &Users{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	// save user to database
}

func validateVariable() {

	myEmail := "joeybloggs.gmail.com"

	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}
