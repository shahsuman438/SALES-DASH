package sales

import "github.com/gin-gonic/gin"

func setupControllers(engine *gin.Engine) {
	engine.Group("/sales")
}
