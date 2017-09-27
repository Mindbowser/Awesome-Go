package main
 	 import (
        		_ "github.com/go-sql-driver/mysql"
        		"database/sql"
        		"log"
    	)
    
    	func main(){
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
		if err != nil {
		log.Fatal(err)
		}
		tx,_:=db.Begin()
		stmt, err := db.Prepare("INSERT INTO user(id,username) VALUES(?,?)")
		res,err:=stmt.Exec(99,"John")
		res,err=stmt.Exec(88,"Martin")
		if err!=nil{
			tx.Rollback()
			log.Fatal(err)
		}
		tx.Commit()
		log.Println(res)		
}
