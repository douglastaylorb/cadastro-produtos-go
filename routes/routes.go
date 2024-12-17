package routes

import (
	"net/http"

	"github.com/douglastaylorb/cadastro-produtos-go/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
