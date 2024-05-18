package utility

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		fmt.Printf("読み込み出来ませんでした: %v", errEnv)
	} 
	user := os.Getenv("AUTH")
	pw := os.Getenv("MYSQL_ROOT_PASSWORD")
	database := os.Getenv("MYSQL_DATABASE")
	var path string = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=True", user, pw, database)
	fmt.Println("path:", path)
	var err error
	if Db, err = sql.Open("mysql", path); err != nil {
		log.Fatal("Db open error:", err.Error())
	}
	checkConnect(100)

	fmt.Println("db connected!!")
}

func checkConnect(count uint) {
	var err error
	if err = Db.Ping(); err != nil {
		time.Sleep(time.Second * 2)
		count--
		fmt.Printf("retry... count:%v\n", count)
		checkConnect(count)
	}
}