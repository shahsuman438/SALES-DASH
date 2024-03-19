package notification

import "github.com/gin-gonic/gin"

func StartModule(engine *gin.Engine) {
	SetupControllers(engine)
}
