package main

import (
	"fmt"
	"github.com/Lyo-Shur/gorm"
	"github.com/Lyo-Shur/gorm/generate/mvc"
	"github.com/Lyo-Shur/gorm/info"
	"log"
)

// 在多数据源的情况下
// 简单使用ROM映射 返回结果为TABLE或MAP
// 优点是不需要建立结构体、不需要手写SQL
// 缺点是只支持普通的增删改查

func main() {
	// 首先初始化数据库连接
	mysqlLink := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=true"
	// 当存在多个数据库时，可以使用别名指定数据库
	alias := "database2"
	gorm.Client.MultiInit(alias, mysqlLink)

	// 获取带别名的数据库信息
	database2 := info.GetMultiDataBase(alias)

	// 获取默认的数据库查询工具 test是表名
	accountService := mvc.GetService(database2, "test")
	// 执行查询
	table, err := accountService.GetList(map[string]string{})
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(table)
	fmt.Println(table.ToMap())
}
