package main

import (
	"fmt"
	"net/http"
)
type Pessoa struct{
    idade int
    nome string
}
func(u Pessoa) ola(){
    fmt.Printf("ola %s", u.nome)
}
func main() {
//    go fmt.Println("Fala, Erick! Projeto Go rodando!")
//      nome := "come"
//   go  fmt.Print(nome)


//     p := Pessoa{
//     idade: 18,
//     nome:"erck",
//      }
// p.ola()
// arraylist := []string{}
// arraylist = append(arraylist, "erick")
// fmt.Print(arraylist[0])

http.HandleFunc("/", handler)
fmt.Print("Servidor iniciado")
http.ListenAndServe(":8000", nil)

}
func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Olá, Erick! Você acessou: %s", r.URL.Path)

}
func soma(a int,  two int) int{
return a+two;
}