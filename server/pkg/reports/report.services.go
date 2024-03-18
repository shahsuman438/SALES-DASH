package reports

import (
	"math"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/product"
	"github.com/shahsuman438/SALES-DASH/CORE-API/pkg/sales"
)

func GetSummeryReports(ctx *gin.Context) (*SummeryReport, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	var mostProfitableProduct int
	var leastProfitableProduct int
	dateOfHighestSales := ""
	dateOfLeastSales := ""
	maxTotalTransactionAmount := 0
	minTotalTransactionAmount := math.MaxInt32

	productTotalTransactionAmounts := make(map[int]int)

	for _, sale := range allSales {
		productTotalTransactionAmounts[sale.ProductId] += sale.TotalTransactionAmount

		if sale.TotalTransactionAmount > maxTotalTransactionAmount {
			maxTotalTransactionAmount = sale.TotalTransactionAmount
			mostProfitableProduct = sale.ProductId
		}

		if sale.TotalTransactionAmount < minTotalTransactionAmount {
			minTotalTransactionAmount = sale.TotalTransactionAmount
			leastProfitableProduct = sale.ProductId
		}
	}

	for _, sale := range allSales {
		if sale.TotalTransactionAmount == maxTotalTransactionAmount {
			dateOfHighestSales = sale.TransactionDate
		}
		if sale.TotalTransactionAmount == minTotalTransactionAmount {
			dateOfLeastSales = sale.TransactionDate
		}
	}
	profitableProduct, err := product.FetchByKeyValue(ctx, "productId", mostProfitableProduct)
	if err != nil {
		return nil, err
	}
	leastProduct, err := product.FetchByKeyValue(ctx, "productId", leastProfitableProduct)
	if err != nil {
		return nil, err
	}
	summaryReport := &SummeryReport{
		MostProfitableProduct:  profitableProduct[0].ProductName,
		LeastProfitableProduct: leastProduct[0].ProductName,
		DateOfHighestSales:     dateOfHighestSales,
		DateOfLeastSales:       dateOfLeastSales,
	}

	return summaryReport, nil

}
