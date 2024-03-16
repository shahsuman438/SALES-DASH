package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/database"
)

func Add(ctx *gin.Context, data *Sales) error {
	err := database.Save(ctx, data, "Sales")
	return err
}

func AddMany(ctx *gin.Context, data *[]Sales) error {
	interfaceData := make([]interface{}, len(*data))
	for i, v := range *data {
		interfaceData[i] = v
	}
	err := database.SaveMany(ctx, interfaceData, "Sales")
	return err
}


