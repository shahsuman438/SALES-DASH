package product

import (
	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/database"
)

func Add(ctx *gin.Context, data *Product) error {
	err := database.Save(ctx, data, "Product")
	return err
}

func AddMany(ctx *gin.Context, data *[]Product) error {
	interfaceData := make([]interface{}, len(*data))
	for i, v := range *data {
		interfaceData[i] = v
	}
	err := database.SaveMany(ctx, interfaceData, "Product")
	return err
}
