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
	maxTotalTransactionAmount := 0.00
	minTotalTransactionAmount := math.MaxFloat64

	productTotalTransactionAmounts := make(map[int]float64)

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

	productMap := make(map[int]product.Product, len(allProduct))
	for _, prdct := range allProduct {
		productMap[prdct.ProductId] = prdct
	}

	salesByProduct := make(map[int]saleByProductItr)
	for _, sale := range allSales {
		if _, ok := salesByProduct[sale.ProductId]; !ok {
			salesByProduct[sale.ProductId] = saleByProductItr{}
		}
		salesData := salesByProduct[sale.ProductId]
		salesData.TotalTransactionAmount += sale.TotalTransactionAmount
		salesData.TotalQuantity += sale.Quantity
		salesByProduct[sale.ProductId] = salesData
	}

	// Create and populate SalesByProduct slice
	var result []SalesByProduct
	var idx = 0
	for key, value := range salesByProduct {
		idx++
		product := productMap[key]
		profit := value.TotalTransactionAmount - (float64(value.TotalQuantity) * product.CostPrice)
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

func GetSalesByBrand(ctx *gin.Context) ([]SalesByBrand, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}
	allProduct, err := product.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	productMap := make(map[int]product.Product, len(allProduct))
	for _, prdct := range allProduct {
		productMap[prdct.ProductId] = prdct
	}

	salesByBrand := make(map[string]saleByBrandItr)
	for _, sale := range allSales {
		pd, ok := productMap[sale.ProductId]
		if !ok {
			continue // Skip if product not found
		}

		brandName := pd.BrandName
		if _, ok := salesByBrand[brandName]; !ok {
			salesByBrand[brandName] = saleByBrandItr{
				MostSoldProduct: make(map[string]int),
			}
		}

		salesData := salesByBrand[brandName]
		salesData.TotalQuantity += sale.Quantity
		salesData.TotalTransactionAmount += sale.TotalTransactionAmount
		salesData.MostSoldProduct[pd.ProductName] += sale.Quantity
		salesData.TotalProfit += calculateProfit(sale, pd)
		salesByBrand[brandName] = salesData
	}

	return convertToSalesByBrandSlice(salesByBrand), nil
}

func calculateProfit(sale sales.Sales, pd product.Product) float64 {
	return float64(sale.TotalTransactionAmount) - (pd.CostPrice * float64(sale.Quantity))
}

func convertToSalesByBrandSlice(salesByBrand map[string]saleByBrandItr) []SalesByBrand {
	var result []SalesByBrand
	idx := 0
	for key, value := range salesByBrand {
		idx++
		profit := value.TotalProfit

		sale := SalesByBrand{
			Sn:                idx,
			BrandName:         key,
			MostSoldProduct:   getMostSoldProduct(value.MostSoldProduct),
			TotalQuantitySold: value.TotalQuantity,
			TotalRevenue:      value.TotalTransactionAmount,
			TotalProfit:       profit,
		}
		result = append(result, sale)
	}
	return result
}

func getMostSoldProduct(value map[string]int) string {
	var maxKey string
	var maxValue int
	for key, value := range value {
		if value > maxValue {
			maxKey = key
			maxValue = value
		}
	}
	return maxKey
}
