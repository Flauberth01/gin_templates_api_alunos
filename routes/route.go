package routes

import (
	"github.com/flauberth01/gin_templates_api_alunos/handlers"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")
	r.GET("/:nome", handlers.Saudacoes)
	r.GET("/alunos", handlers.TodosAlunos)
	r.GET("/alunos/:id", handlers.BuscarAlunoPorID)
	r.POST("/alunos", handlers.CriarNovoAluno)
	r.DELETE("/alunos/:id", handlers.DeletarAluno)
	r.PATCH("/alunos/:id", handlers.EditarAluno)
	r.GET("/alunos/cpf/:cpf", handlers.BuscaAlunoPorCPF)
	r.GET("/alunos/", handlers.BuscaAlunoPorCPF)
	r.GET("/index", handlers.ExibePaginaIndex)
	r.GET("/detalhes/alunos/:id", handlers.ExibePaginaDetalhes)
	r.Run()
}
