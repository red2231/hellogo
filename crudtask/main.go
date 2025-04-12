package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	 db, err:=sql.Open("mysql", "root:erick@unix(/var/run/mysqld/mysqld.sock)/golang?parseTime=true")
	 if err!=nil{
		log.Fatal(err.Error())
		return
	 }
mutex:=NewServer(db)
defer db.Close()
fmt.Print("Servidor iniciado na porta 4000")
http.ListenAndServe(":4000", mutex)
}