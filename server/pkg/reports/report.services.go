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

func GetSalesByProduct(ctx *gin.Context) ([]SalesByProduct, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	allProduct, err := product.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	productMap := make(map[int]product.Product)

	salesByProduct := make(map[int]saleByProductItr)

	// Aggregate sales data by product
	for _, sale := range allSales {
		if _, ok := salesByProduct[sale.ProductId]; !ok {
			salesByProduct[sale.ProductId] = saleByProductItr{}
		}
		salesData := salesByProduct[sale.ProductId]
		salesData.TotalTransactionAmount += sale.TotalTransactionAmount
		salesData.TotalQuantity += sale.Quantity
		salesByProduct[sale.ProductId] = salesData
	}

	for _, prdct := range allProduct {
		if _, ok := productMap[prdct.ProductId]; !ok {
			productMap[prdct.ProductId] = product.Product{}
		}
		productMap[prdct.ProductId] = prdct
	}

	// Create and populate SalesByProduct slice
	var result []SalesByProduct
	var idx = 0
	for key, value := range salesByProduct {
		idx += 1
		product := productMap[key]
		profit := value.TotalTransactionAmount - (value.TotalQuantity * product.CostPrice)
		sale := SalesByProduct{
			Sn:                idx,
			ProductName:       product.ProductName,
			BrandName:         product.BrandName,
			Category:          product.Category,
			TotalQuantitySold: value.TotalQuantity,
			TotalRevenue:      value.TotalTransactionAmount,
			TotalProfit:       profit,
		}
		result = append(result, sale)
	}

	return result, nil
}
