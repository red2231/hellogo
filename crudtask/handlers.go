package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)



func(s *Server) getId(w http.ResponseWriter, r *http.Request){

	id, _:=strconv.Atoi(r.PathValue("id")) 

var task Task

sql:= "select name, description, scheduled from task where id = ?"
result:=s.db.QueryRow(sql, id)
err := result.Scan(&task.Name, &task.Description, &task.Scheduled)
if err!=nil{
log.Print("Nada encontrado")
http.Error(w, "Nenhum resultado encontrado", http.StatusNotFound)
return
}
value, err:=json.Marshal(task)
if err!=nil{
	log.Fatal(err.Error())
}
io.Writer.Write(w, value)
}
func(s *Server) getAll(w http.ResponseWriter, r *http.Request){

}