package routes

import (
	controllers "github.com/LandimTiago/gin-api-golang/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()

	r.GET("/alunos", controllers.RetornaAlunos)

	r.GET("/alunos/:nome", controllers.RetornaSaudacao)

	r.Run()
}
