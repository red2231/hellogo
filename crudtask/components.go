package main

import (
	"database/sql"
	"net/http"
	"time"
)

type Task struct{
	Name string `json:"name"`
	Description string `json:"description"`
	Scheduled time.Time `json:"scheduled"`
}
type Server struct{
	db *sql.DB
	mux *http.ServeMux
}
func NewServer(db *sql.DB)*Server{
s:=&Server{db: db, mux: http.NewServeMux()}
s.routes()
return s
}

func(s *Server) routes(){
s.mux.HandleFunc("GET /task/{id}", s.getId)
s.mux.HandleFunc("GET /task", s.getAll)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request){
	s.mux.ServeHTTP(w, r)
}

