package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/my/repo/models"
	"log"
	"strings"
	"time"
)

//定义结构体(xorm支持双向映射)
type User struct {
	Id              int64     `xorm:"pk autoincr"` //指定主键并自增
	Name            string    `xorm:"name"`
	Real_name       string    `xorm:"real_name"`
	Age             int       `xorm:"age"`
	Sex             int       `xorm:"sex"`
	Passwrod        string    `xorm:"password"`
	Mobile          string    `xorm:"mobile"`
	Created_at      time.Time `xorm:"created_at"`
	Last_login_time time.Time `xorm:"last_login_time"`
}

//定义orm引擎，手册地址：http://books.studygolang.com/xorm/chapter-05/10.exist.html
var db *xorm.Engine

func main() {
	var err error
	db, err = xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := db.Sync(new(User)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}
	defer db.Close()

	// sql分段拼接用法
	var d []models.Rows
	var dbSql strings.Builder
	params := make([]interface{}, 0)
	// 方法一(推荐)
	dba := db.Table("rows").Where("a=?", 1)
	dba.Where("b=?", 1)
	dba.Or("c>=?", 1)
	dba.Find(&d)
	fmt.Println(d)
	fmt.Println(dba.LastSQL())

	// 方法二
	dbSql.WriteString("select * from rows where a = ?")
	params = append(params, 1)
	dbSql.WriteString(" and b = ? ")
	params = append(params, 1)
	db.SQL(dbSql.String(), params...).Find(&d)
	fmt.Println(d)

	// 表映射
	var data models.Rows
	data.E = 456
	after := func(bean interface{}) {
		fmt.Println("after", bean)
	}
	// 事件处理 参考文档：http://books.studygolang.com/xorm/chapter-12/
	db.Before(after).ID(3).Update(&data)

	// 执行sql命令
	sql := "update `rows` set c=? where id=?"
	res, err := db.Exec(sql, 12, 1)
	fmt.Println(res)

	// 聚合方法
	total, err := db.Where("a=?", 1).Count(&data)
	total2, err := db.Where("a=?", 1).Sum(data, "b")               // Sum 求某个字段的和，返回float6
	total3, err := db.Where("a=?", 1).SumInt(data, "b")            // SumInt 求某个字段的和，返回int64
	total4, err := db.Where("a=?", 1).Sums(data, "b", "c", "d")    // Sums 求某几个字段的和， 返回float64的Slice
	total5, err := db.Where("a=?", 1).SumsInt(data, "b", "c", "d") // Sums 求某几个字段的和， 返回float64的Slice
	if err != nil {

	}
	fmt.Println(total, total2, total3, total4, total5)
	fmt.Println(data)
	return
}

/**
如果你的需求是：判断某条记录是否存在，若存在，则返回这条记录。
建议直接使用Get方法。
如果仅仅判断某条记录是否存在，则使用Exist方法，Exist的执行效率要比Get更高。
*/
func exsit() {
	var data models.Rows
	has, err := db.Where("a=?", 2).Exist(&data)
	if err != nil {
		return
	}
	fmt.Println(has)
}

// 事务
func transaction() {
	var data models.Rows
	session := db.NewSession()
	defer session.Close()

	session.Begin()
	data.E = 11
	affect, err := session.ID(3).Update(&data)
	if err != nil {
		session.Rollback()
		return
	}
	fmt.Println(affect)
	session.Commit()
}

//查 参考文档：http://books.studygolang.com/xorm/chapter-05/1.conditions.html
func get(id int64) {
	// 查询条件  limit注意值的顺序
	db.Table("rows").Alias("r").Distinct("a").Cols("id").In("a", []int{1, 2, 3}).Where("a=?",
		1).And("b=?", 1).OrderBy("id desc").Desc("id").Asc("a").GroupBy("a").Having("total>0").Limit(10, 0)

	// like用法 http://books.studygolang.com/xorm/chapter-14/
	name := "test"
	db.Where("a like ?", "%"+name+"%")

	var res []models.Rows
	db.Select("a.*, (SELECT b FROM `rows` limit 1) as name").Find(&res)
	db.SQL("SELECT * FROM `rows`").Find(&res)

	// 单记录
	data := &models.Rows{Id: 1}
	db.ID(1).Get(&data)
	is, _ := db.Get(data)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	// 多记录
	more := make([]*models.Rows, 0)
	db.Find(&more)

}

//改
func update(id int64, user *User) bool {
	affected, err := db.ID(id).Update(user)
	if err != nil {
		log.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}

//增 http://books.studygolang.com/xorm/chapter-04/
func insert() () {
	// 单条插入
	data := new(models.Rows)
	data.B = 123
	affected, err := db.Insert(data)

	/**
	插入多条记录并且不使用批量插入，此时实际生成多条插入语句，每条记录均会自动赋予Id值
	这里虽然支持同时插入，但这些插入并没有事务关系。因此有可能在中间插入出错后，后面的插入将不会继续。此时前面的插入已经成功，如果需要回滚，请开启事务
	限制，因此这样的语句有一个最大的记录数，根据经验测算在150条左右。大于150条后，生成的sql语句将太长可能导致执行失败。因此在插入大量数据时，目前需要自行分割成每150条插入一次。
	*/
	users := make([]*models.Rows, 1)
	users[0] = new(models.Rows)
	users[0].C = 12
	affected2, err := db.Insert(&users)
	if err != nil {

	}
	fmt.Println(affected, affected2)
}

//删
func del() {
	data := new(models.Rows)
	// Deleted可以让您不真正的删除数据，而是标记一个删除时间。使用此特性需要在xorm标记中使用deleted标记，如下所示进行标记，对应的字段必须为time.Time类型。
	db.Id(1).Delete(&data)
	// 那么如果记录已经被标记为删除后，要真正的获得该条记录或者真正的删除该条记录，需要启用Unscoped
	var row models.Rows
	// 此时将可以获得删除记录
	db.Id(1).Unscoped().Get(&row)
	// 此时将可以真正的删除记录
	db.Id(1).Unscoped().Delete(&row)
}
