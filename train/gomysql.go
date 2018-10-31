package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"reflect"
)

func main() {
	db, err := sql.Open("mysql","root:123456@tcp(namenode:3306)/test?charset=utf8")
	if err != nil{
		fmt.Println(err)
	}
	rows, _ :=db.Query("select * from  detail")
	fmt.Println(db)
	fmt.Println(rows)
	defer rows.Close()
	//ra,_ := rows.Columns()
	//fmt.Println(ra)
	//for i :=range ra{
	//	fmt.Println(ra[i])
	//}
	for rows.Next(){
		var col1 string
		var col2 float64
		fmt.Println(reflect.TypeOf(rows))
		err = rows.Scan(&col1,&col2)
		fmt.Println(col1,col2)
	}
}
