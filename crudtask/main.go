package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	chanel:=make(chan *sql.DB )
	erru:= make(chan error)

	go func(){
canal, err:=sql.Open("mysql", "root:erick@unix(/var/run/mysqld/mysqld.sock)/golang?parseTime=true")
if err!=nil{
	
	erru<-err
	return
}
chanel<-canal
	}()

select{
	case db:=<-chanel:
mutex:=NewServer(db)
defer db.Close()
fmt.Print("Servidor iniciado na porta 4000")
 http.ListenAndServe(":4000", mutex)
 case err:=<-erru:
		log.Fatal(err.Error())
		return
	 
}
	

}