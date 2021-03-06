package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"        //
	_ "github.com/jinzhu/gorm/dialects/mysql" //
)

// DB2 适配其他版本的数据库连接实例
var DB2 *sql.DB

//  创建数据库连接实例
func Database(conString string, conString2 string) {
	db2, err := sql.Open("mysql", conString2)

	if err != nil {
		panic("fail to connect database2")
	}

	DB2 = db2
}
