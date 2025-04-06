package main

import (
	"fmt"
	"net/http"
)
type Pessoa struct{
    idade int
    nome string
}
type animalInterface interface{
    falar() string
}
func (u Pessoa) falar() string{
    return "ola mundo"
}
func(u Pessoa) ola(){
    fmt.Printf("ola %s", u.nome)
}

func ola(a animalInterface){
    fmt.Print(a.falar())
}
func main() {
//    go fmt.Println("Fala, Erick! Projeto Go rodando!")
//      nome := "come"
//   go  fmt.Print(nome)


    p := Pessoa{
    idade: 18,
    nome:"erck",
     }
p.ola() 
ola(p)
// arraylist := []string{}
// arraylist = append(arraylist, "erick")
// fmt.Print(arraylist[0])

// http.HandleFunc("/", handler)
// fmt.Print("Servidor iniciado")
// http.ListenAndServe(":8000", nil)
// resp, err := http.Get("https://www.redhat.com/pt-br/topics/api/what-are-application-programming-interfaces")

// if err !=nil{
//     fmt.Print("erro na requisicao")

// }
// defer resp.Body.Close()
// fmt.Print(resp.Header.Values("Content-Type"))
// }

http.HandleFunc("POST /ola", handler)

http.ListenAndServe(":5000", nil)
}

func handler(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Olá, Erick! Você acessou: %s", r.URL.Path)

}
func soma(a int,  two int) int{
return a+two;
}