package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)



func(s *Server) getId(w http.ResponseWriter, r *http.Request){
	id, _:=strconv.Atoi(r.PathValue("id")) 
var task Task
erro:=make(chan error)
sq:= "select name, description, scheduled from task where id = ?"
rows:= make(chan *sql.Row)
go func(){
	rows<-s.db.QueryRow(sq, id)
}()
result:=<-rows
go func(){ erro<-result.Scan(&task.Name, &task.Description, &task.Scheduled)}()
err:=<-erro
if err!=nil{
log.Print("Nada encontrado")
http.Error(w, "Nenhum resultado encontrado", http.StatusNotFound)
return
}
errChan := make(chan error, 1)
byteChan := make(chan []byte, 1)

go func() {
    value, err := json.Marshal(task)
    if err != nil {
        errChan <- err
        return
    }
    byteChan <- value
}()

select {
case err := <-errChan:
    if err != nil {
        log.Fatal(err)
        return
    }
case bit := <-byteChan:
   io.Writer.Write(w, bit)
}}
func(s *Server) getAll(w http.ResponseWriter, r *http.Request) {
all:="select name, description, scheduled from task"
values, err:=s.db.Query(all)
if err!=nil{
	log.Fatal(err)
}
 tasks := []Task{}
for values.Next(){
task:=new(Task)
values.Scan(&task.Name, &task.Description, &task.Scheduled)
tasks = append(tasks, *task)
}
val, err:=json.Marshal(tasks)
if err!=nil{log.Fatal(err)
return}

io.Writer.Write(w, val)
}
func(s *Server) createTask(w http.ResponseWriter, r *http.Request){
var task Task
bites, err:=io.ReadAll(r.Body)
if err!=nil{
	log.Fatalf("erro na leitura %s", err)
}
erro:=json.Unmarshal(bites, &task)
if erro!=nil{
	log.Fatal(erro.Error())
}

go s.db.Exec("insert into task(name, description, scheduled) values (?,?,?)", task.Name,task.Description,task.Scheduled)
 w.WriteHeader(http.StatusCreated)
io.WriteString(w, "tarefa criada com sucesso")
}