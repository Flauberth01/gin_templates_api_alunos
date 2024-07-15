package main

import (
	"github.com/flauberth01/gin_templates_api_alunos/database"
	"github.com/flauberth01/gin_templates_api_alunos/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
