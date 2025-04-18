package main

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)


func reqToTask(r http.Request)*Task{
var task Task
bites, err:=io.ReadAll(r.Body)
if err!=nil{
	log.Fatal(err.Error())
return nil
}

err=json.Unmarshal(bites, &task)
if err!=nil{
	log.Fatal(err.Error())
	return nil
}

return &task
}
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
	var task Task

values.Scan(&task.Name, &task.Description, &task.Scheduled)
tasks = append(tasks, task)
}
val, err:=json.Marshal(tasks)
if err!=nil{log.Fatal(err)
return}
io.Writer.Write(w, val)
}
func(s *Server) createTask(w http.ResponseWriter, r *http.Request){
task:=reqToTask(*r)
result:=make(chan sql.Result)
er:=make(chan error)
go func(){
 value, err:=s.db.Exec("insert into task(name, description, scheduled) values (?,?,?)", task.Name,task.Description,task.Scheduled)
if err!=nil{
	er<-err
	return
}
result<-value
}()
select{
	case ero:= <-er:
	log.Fatal(ero.Error())
	return
	case <-result:
 w.WriteHeader(http.StatusCreated)
io.WriteString(w, "tarefa criada com sucesso")
}}
func(s *Server) updateTask(w http.ResponseWriter, r *http.Request){
task :=reqToTask(*r)
go s.db.Exec("UPDATE task SET name=?, description=?, scheduled =? where id =?", task.Name, task.Description, task.Scheduled, task.Id)
w.WriteHeader(http.StatusOK)
io.WriteString(w, "tarefa atualizada")
}
func(s *Server) deleteTask(w http.ResponseWriter, r *http.Request){
		id, _:=strconv.Atoi(r.PathValue("id")) 
		value :=make(chan sql.Result)
		erro:=make(chan error)
		go func(){
valuee, err:=s.db.Exec("Delete from task where id =?", id)
		if err!=nil{
			
			erro<-err
			return
		}
value<-valuee
		}()
	select {
	
	case erra:=<-erro:
	
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Tarefa não encontrada para deleção")
	log.Println(erra.Error())
		return
	case <-value:
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "tarefa deletadas")
	return
	}

}