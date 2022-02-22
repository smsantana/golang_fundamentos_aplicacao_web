package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul Bem Bonita", Preco: 100, Quantidade: 10},
		{"Notebook", "Muito Rapido", 1999, 2},
		{"Fone", "Muito Bom", 100, 4},
		{"Oculos", "Escuro", 50.0, 6},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
