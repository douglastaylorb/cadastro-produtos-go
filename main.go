package main

import (
	"net/http"

	"github.com/douglastaylorb/cadastro-produtos-go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
