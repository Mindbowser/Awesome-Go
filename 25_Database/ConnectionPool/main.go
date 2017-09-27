package main
 	 import (
        		_ "github.com/go-sql-driver/mysql"
        		"database/sql"
        		"log"
    	)
    
    	func main(){
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/employeedb")
		
		// Connection Pooling methods
		
		db.SetConnMaxLifetime(500)
		db.SetMaxIdleConns(50)
		db.SetMaxOpenConns(10)
		db.Stats()
		
		if err != nil {
		log.Fatal(err)
		}
		tx,_:=db.Begin()
		stmt, err := tx.Prepare("INSERT INTO user(id,username) VALUES(?,?)")
		res,err:=stmt.Exec(4,"Abhijit")
		res,err=stmt.Exec(5,"Yogesh")
		if err!=nil{
			tx.Rollback()
			log.Fatal(err)
		}
		tx.Commit()
		log.Println(res)
		
}
