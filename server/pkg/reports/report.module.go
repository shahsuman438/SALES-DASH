package reports

import "github.com/gin-gonic/gin"

func StartModule(engine *gin.Engine) {
	setupControllers(engine)
}
