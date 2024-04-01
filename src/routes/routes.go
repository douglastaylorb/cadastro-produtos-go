package routes

import (
	"net/http"
	"src/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
