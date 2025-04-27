package router

import (
	"desafio-itau-golang/controller"
	"github.com/gin-gonic/gin"
)

func HandleRequest(r *gin.Engine) {
	router := r.Group("/")

	router.POST("transacao", controller.CreateTransaction)
	// request.GET("estatisticas", controller.GetStatistics)
	router.DELETE("transacao/:id", controller.DeleteTransaction)
}
