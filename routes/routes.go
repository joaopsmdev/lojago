package routes

import (
	"Projetos_GO/Loja/controllers"
	"net/http"
)

func CarregaRotas() {

	http.HandleFunc("/", controllers.Index) //atende requisição
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.ListenAndServe(":8000", nil)
}
