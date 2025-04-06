package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)


func main(){
	
db, err := sql.Open("mysql","root:erick@tcp(127.0.0.1:3306)/bot")
if err !=nil{
fmt.Print("Erro ao conectar")
}
defer db.Close()
conecto :=db.Ping()
if conecto !=nil{
log.Fatal("erro")
}

rows, err :=db.Query("select username, id from usuario")
if err!=nil{
	log.Fatal("erro")
}
defer rows.Close()
for rows.Next(){
	var username string 
	var id int
	err = rows.Scan(&username, &id)
	if err !=nil{
		log.Fatal("nao encontrado")
	}
	fmt.Printf("Username: %s\n e id: %d\n", username, id)
}
}