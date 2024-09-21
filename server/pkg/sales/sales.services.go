package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
	"github.com/shahsuman438/SALES-DASH/server/pkg/database"
	"github.com/shahsuman438/SALES-DASH/server/pkg/utils/logger"
)

var collectionName = "Sales"

func Add(ctx *gin.Context, data *Sales) error {
	err := database.Save(ctx, data, collectionName)
	return err
}

func AddMany(ctx *gin.Context, data *[]Sales) error {
	interfaceData := make([]interface{}, len(*data))
	for i, v := range *data {
		interfaceData[i] = v
	}
	err := database.SaveMany(ctx, interfaceData, collectionName)
	return err
}

func Fetch(ctx *gin.Context) ([]Sales, error) {
	data, err := database.Fetch(ctx, collectionName)
	if err != nil {
		logger.Error("Error fetching from Product", err)
		return nil, err
	}
	var transactions []Sales
	for _, item := range data {
		var sale Sales
		if err := mapstructure.Decode(item, &sale); err != nil {
			logger.Error("Error decoding product", err)
			return nil, err
		}
		transactions = append(transactions, sale)
	}
	return transactions, nil
}
