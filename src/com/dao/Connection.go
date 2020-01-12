package dao

import (
	"com/util"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	NAME string
}

var dbLogPath = "./db_log.log"
var connection = connect()

func connect() *sql.DB {
	user := "root"
	password := ""
	database := "test"
	host := "127.0.0.1"
	port := "3306"

	db, err := sql.Open("mysql", user+":"+password+"@tcp("+host+":"+port+")/"+database)

	if err != nil {
		util.LogToFile(dbLogPath, err.Error())
	}

	return db
}

func GetUsers() [] User{

	rows, e := connection.Query("SELECT ID, name FROM user")

	if e != nil {
		util.LogToFile(dbLogPath, e.Error())
	}


	var s [] User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.NAME)

		if err != nil {
			util.LogToFile(dbLogPath,err.Error())
		}
		s = append(s, user)
	}

	defer connection.Close()

	return s

}
