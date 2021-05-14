package service

import (
	"capybara/db/mysql"
)



func ManagerLogin(username string, password string) (err error) {
	db := mysql.MysqlDB.GetConn()
	//查询数据库中是否存在此条数据
	db.Where("username=?", username).First().Error
}
