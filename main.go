package main

import (
	"Projetos_GO/Loja/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
}
