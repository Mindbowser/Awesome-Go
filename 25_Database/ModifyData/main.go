package main
 	 import (
        		_ "github.com/go-sql-driver/mysql"
        		"database/sql"
        		"fmt"
        		"log"
    	)
    
    	func main(){
    		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
		if err != nil {
		log.Fatal(err)
		}else{
			fmt.Println("Connection Established")
		}
		stmt, err := db.Prepare("INSERT INTO user(id,username) VALUES(?,?)")
		if err != nil {
			log.Fatal(err)
		}
		res, err := stmt.Exec(3,"Simon")
		if err != nil {
			log.Fatal(err)
		}
		lastId, err := res.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		rowCnt, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		
		log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
}
