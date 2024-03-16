package file

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/product"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/utils/logger"
)

func ProcessProductFiles(pathToCsv string) error {
	file, err := os.Open(pathToCsv)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	if _, err := reader.Read(); err != nil {
		logger.Error("Error reading header", err)
		return err
	}

	var products []product.Product
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error("Error reading record:", err)
			return err
		}
		var product product.Product
		pId, err := strconv.Atoi(record[0])
		if err != nil {
			logger.Error("Error Converting productId", err)
			continue
		}
		product.ProductId = pId
		product.ProductName = record[1]
		product.BrandName = record[2]

		cp, err := strconv.Atoi(record[3])
		if err != nil {
			logger.Error("Error Converting Cost_Price", err)
			continue
		}
		product.CostPrice = cp

		sp, err := strconv.Atoi(record[4])
		if err != nil {
			logger.Error("Error Converting Cost_Price", err)
			continue
		}
		product.SellingPrice = sp

		product.Category = record[5]
		product.ExpiryDate = record[6]

		products = append(products, product)
	}

	err = product.AddMany(&gin.Context{}, &products)
	if err != nil {
		logger.Error("Unable to add data in Produc", err)
	}
	logger.Info(fmt.Sprintf("PROCESSED File %s ", pathToCsv))
	return nil
}
