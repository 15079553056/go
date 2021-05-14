package mysql

import (
	"log"

	"github.com/jinzhu/gorm"

	//这一行需要保留，否则会报import _ "github.com/go-sql-driver/mysql"错误
	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *Mysql

type Mysql struct {
	conn *gorm.DB
}

func Default() {
	mysql := &Mysql{}
	conn, err := gorm.Open("mysql", "root:root@(122.51.138.213:3306)/db_capybara?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err.Error())
		panic(err.Error())
	}

	//设置数据库连接池最大连接数
	conn.DB().SetMaxOpenConns(100)
	//连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭
	conn.DB().SetMaxIdleConns(20)
	mysql.conn=conn
	MysqlDB=mysql
}

func (mysql *Mysql) GetConn() *gorm.DB{
	return mysql.conn
}