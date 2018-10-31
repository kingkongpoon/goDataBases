package main

import (
	"database/sql"
	"fmt"
	_ "github.com/prestodb/presto-go-client/presto"
	"log"
)

func main() {

	dsn := "http://hive@bdmaster1:9999?catalog=hive&schema=default"
	db, _ := sql.Open("presto", dsn)
	//if err != nil{
	//	fmt.Println(err)
	//}
	rows, err :=db.Query("select * from t_10001")
	if err != nil{
		fmt.Println(err)
	}

	defer rows.Close()

	cols , _ := rows.Columns()
	fmt.Println(cols)


	values := make([]sql.RawBytes, len(cols))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Print(cols[i], ": ", value,",")
		}
		fmt.Println("------------------")
	}
}
