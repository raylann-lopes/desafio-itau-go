package router

import (
	"desafio-itau-golang/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest(r *gin.Engine) {
	request := r.Group("/")

	request.POST("transacao", controller.CreateTransaction)
	// request.GET("estatisticas", controller.GetStatistics)
	// request.DELETE("transaction/:id", controller.DeleteTransaction)
}
