package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// MySQL 连接信息
	dataSourceName := "root:P@ssword1@tcp(192.168.0.214:21406)/"

	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//// 检查连接是否正常
	//err = db.Ping()
	//if err != nil {
	//	log.Fatal("Error connecting to the database: ", err)
	//}
	//fmt.Println("Connected to MySQL database!")

	// 查询所有数据库
	rows, err := db.Query("SHOW DATABASES")
	if err != nil {
		log.Fatal("Error fetching databases: ", err)
	}
	defer rows.Close()

	// 遍历查询结果
	var databaseName string
	fmt.Println("List of databases:")
	for rows.Next() {
		err := rows.Scan(&databaseName)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(databaseName)
	}
	//// 检查遍历时是否有错误
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
}
