package common

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:q1w2e3r4@tcp(192.168.31.139:3306)/hl?charset=utf8&parseTime=True")
	if nil != err {
		fmt.Println(err)
	}

	fmt.Println("数据库初始化成功")
}
