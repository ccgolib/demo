package main

import (
	"fmt"
	"os"
	"time"
)

func main()  {
	//f, err := os.Create("1.sql")
	f, err := os.OpenFile("1.sql", os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}else {
		// 生成百万数据sql
		for i:= 0; i<1000000; i++ {
			name := fmt.Sprintf("昵称No%d", i)
			real_name := fmt.Sprintf("真实名%d", i)
			age := i
			sex := 1
			password := fmt.Sprintf("pwd%010d", i)
			mobile := fmt.Sprintf("138%08d", i)
			created_at := time.Now().Format("2006-01-02 15:04:05")
			last_login_time := time.Now().Format("2006-01-02 15:04:05")

			insert := fmt.Sprintf("INSERT INTO `user` VALUES (%d, '%s', '%s', %d, %d, '%s', '%s', '%s', '%s');\t\n",i, name, real_name,age,sex,password,mobile,created_at,last_login_time)
			n, _ := f.Seek(0, 2)
			_, err = f.WriteAt([]byte(insert), n)
		}
	}

}
