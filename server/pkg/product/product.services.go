package product

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/database"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

var collectionName = "Product"

func Add(ctx *gin.Context, data *Product) error {
	err := database.Save(ctx, data, collectionName)
	return err
}

func AddMany(ctx *gin.Context, data *[]Product) error {
	interfaceData := make([]interface{}, len(*data))
	for i, v := range *data {
		interfaceData[i] = v
	}
	err := database.SaveMany(ctx, interfaceData, collectionName)
	return err
}

func Fetch(ctx *gin.Context) ([]Product, error) {
	data, err := database.Fetch(ctx, collectionName)
	if err != nil {
		logger.Error("Error fetching from Product", err)
		return nil, err
	}
	var products []Product
	for _, item := range data {
		var product Product
		if err := mapstructure.Decode(item, &product); err != nil {
			logger.Error("Error decoding product", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func FetchByKeyValue(ctx *gin.Context, key string, value interface{}) ([]Product, error) {
	data, err := database.FetchByKeyValue(ctx, collectionName, key, value)
	if err != nil {
		return nil, err
	}
	fmt.Println(data)
	var products []Product
	for _, item := range data {
		var product Product
		if err := mapstructure.Decode(item, &product); err != nil {
			logger.Error("Error decoding product", err)
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}
