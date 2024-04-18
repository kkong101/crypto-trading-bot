package webhook

import gin "github.com/gin-gonic/gin"

// NewRouter gin router 생성
func NewRouter() *gin.Engine {

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	router.POST("/kkong101/webhook", WebhookHandler)

	return router
}
