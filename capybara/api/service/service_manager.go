package service

import (
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
)

func ManagerLogin(context *gin.Context) {
	//连接数据库
	connstr := "root:root@tcp(127.0.0.1:3306)/db_capybara"
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	//创建数据库表
	_,err=db.Exec("create table person("+
	"name varchar(12) not null,"+
	"age int default 1"+
	");")

	if err!=nil{
		log.Fatal(err.Error())
		return
	}
}
