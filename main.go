package main

import (
	"net/http"
	"text/template"
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
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
		{"Camiseta", "Camiseta preta", 39.90, 10},
		{"Calça", "Calça jeans", 79.90, 5},
		{"Tênis", "Tênis esportivo", 129.90, 3},
		{"Boné", "Boné preto", 19.90, 20},
	}

	temp.ExecuteTemplate(w, "Index", produtos)
}
