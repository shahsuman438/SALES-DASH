package reports

import (
	"math"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/shahsuman438/SALES-DASH/server/pkg/product"
	"github.com/shahsuman438/SALES-DASH/server/pkg/sales"
)

// Struct to hold aggregation of sale values
type SaleAggregation struct {
	TotalTransactionAmount float64
	TotalQuantity          int
}

func GetSalesByProduct(ctx *gin.Context) ([]SalesByProduct, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	allProducts, err := product.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	productMap := make(map[int]product.Product, len(allProducts))
	for _, prdct := range allProducts {
		productMap[prdct.ProductId] = prdct
	}

	salesByProduct := make(map[int]SaleAggregation)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	numGoroutines := 4
	chunkSize := len(allSales) / numGoroutines

	// Process sales in parallel
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(allSales)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			localSalesByProduct := make(map[int]SaleAggregation)

			for j := start; j < end; j++ {
				sale := allSales[j]
				data := localSalesByProduct[sale.ProductId]
				data.TotalTransactionAmount += sale.TotalTransactionAmount
				data.TotalQuantity += sale.Quantity
				localSalesByProduct[sale.ProductId] = data
			}

			mu.Lock()
			for k, v := range localSalesByProduct {
				data := salesByProduct[k]
				data.TotalTransactionAmount += v.TotalTransactionAmount
				data.TotalQuantity += v.TotalQuantity
				salesByProduct[k] = data
			}
			mu.Unlock()
		}(start, end)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Generate the result
	result := make([]SalesByProduct, 0, len(salesByProduct))
	for productId, salesData := range salesByProduct {
		product := productMap[productId]
		profit := salesData.TotalTransactionAmount - float64(salesData.TotalQuantity)*product.CostPrice
		result = append(result, SalesByProduct{
			ProductName:       product.ProductName,
			BrandName:         product.BrandName,
			Category:          product.Category,
			TotalQuantitySold: salesData.TotalQuantity,
			TotalRevenue:      salesData.TotalTransactionAmount,
			TotalProfit:       profit,
		})
	}

	return result, nil
}

// Summarized sales report with parallel processing
func GetSummeryReports(ctx *gin.Context) (*SummeryReport, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	var (
		mostProfitableProduct   int
		leastProfitableProduct  int
		dateOfHighestSales      string
		dateOfLeastSales        string
		maxTotalTransaction     = 0.00
		minTotalTransaction     = math.MaxFloat64
		productTotalTransaction = make(map[int]float64)
		mu                      sync.Mutex
		wg                      sync.WaitGroup
		numGoroutines           = 4
		chunkSize               = len(allSales) / numGoroutines
	)

	// Process sales data in parallel
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(allSales)
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			localMaxTransaction := 0.00
			localMinTransaction := math.MaxFloat64
			localMaxProduct := 0
			localMinProduct := 0
			localDateMax := ""
			localDateMin := ""

			localTotals := make(map[int]float64)

			for j := start; j < end; j++ {
				sale := allSales[j]
				localTotals[sale.ProductId] += sale.TotalTransactionAmount

				if localTotals[sale.ProductId] > localMaxTransaction {
					localMaxTransaction = localTotals[sale.ProductId]
					localMaxProduct = sale.ProductId
					localDateMax = sale.TransactionDate
				}

				if localTotals[sale.ProductId] < localMinTransaction {
					localMinTransaction = localTotals[sale.ProductId]
					localMinProduct = sale.ProductId
					localDateMin = sale.TransactionDate
				}
			}

			mu.Lock()
			for k, v := range localTotals {
				productTotalTransaction[k] += v
			}
			if localMaxTransaction > maxTotalTransaction {
				maxTotalTransaction = localMaxTransaction
				mostProfitableProduct = localMaxProduct
				dateOfHighestSales = localDateMax
			}
			if localMinTransaction < minTotalTransaction {
				minTotalTransaction = localMinTransaction
				leastProfitableProduct = localMinProduct
				dateOfLeastSales = localDateMin
			}
			mu.Unlock()

		}(start, end)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Fetch the product names for the summary
	mostProduct, err := product.FetchByKeyValue(ctx, "productId", mostProfitableProduct)
	if err != nil {
		return nil, err
	}
	leastProduct, err := product.FetchByKeyValue(ctx, "productId", leastProfitableProduct)
	if err != nil {
		return nil, err
	}

	return &SummeryReport{
		MostProfitableProduct:  mostProduct[0].ProductName,
		LeastProfitableProduct: leastProduct[0].ProductName,
		DateOfHighestSales:     dateOfHighestSales,
		DateOfLeastSales:       dateOfLeastSales,
	}, nil
}

func GetSalesByBrand(ctx *gin.Context) ([]SalesByBrand, error) {
	allSales, err := sales.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	allProducts, err := product.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	// Map products by ID
	productMap := make(map[int]product.Product, len(allProducts))
	for _, prdct := range allProducts {
		productMap[prdct.ProductId] = prdct
	}

	// Aggregate sales by brand using concurrent goroutines
	salesByBrand := make(map[string]saleByBrandItr)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	numGoroutines := 4
	chunkSize := len(allSales) / numGoroutines

	// Process sales in parallel
	for i := 0; i < numGoroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numGoroutines-1 {
			end = len(allSales) // Handle remainder
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()

			localSalesByBrand := make(map[string]saleByBrandItr)

			for j := start; j < end; j++ {
				sale := allSales[j]
				product, ok := productMap[sale.ProductId]
				if !ok {
					continue // Skip invalid products
				}

				brandName := product.BrandName
				data := localSalesByBrand[brandName]

				// Initialize MostSoldProduct map if nil
				if data.MostSoldProduct == nil {
					data.MostSoldProduct = make(map[string]int)
				}

				data.TotalQuantity += sale.Quantity
				data.TotalTransactionAmount += sale.TotalTransactionAmount
				data.MostSoldProduct[product.ProductName] += sale.Quantity
				data.TotalProfit += calculateProfit(sale, product)
				localSalesByBrand[brandName] = data
			}

			// Merge local results into shared salesByBrand map
			mu.Lock()
			for brand, data := range localSalesByBrand {
				globalData := salesByBrand[brand]

				// Initialize MostSoldProduct map in global data if nil
				if globalData.MostSoldProduct == nil {
					globalData.MostSoldProduct = make(map[string]int)
				}

				globalData.TotalQuantity += data.TotalQuantity
				globalData.TotalTransactionAmount += data.TotalTransactionAmount
				globalData.TotalProfit += data.TotalProfit
				for productName, quantity := range data.MostSoldProduct {
					globalData.MostSoldProduct[productName] += quantity
				}
				salesByBrand[brand] = globalData
			}
			mu.Unlock()

		}(start, end)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	// Convert aggregated data to slice
	return convertToSalesByBrandSlice(salesByBrand), nil
}

// Helper to calculate profit
func calculateProfit(sale sales.Sales, product product.Product) float64 {
	return sale.TotalTransactionAmount - (product.CostPrice * float64(sale.Quantity))
}

// Helper to convert map to slice
func convertToSalesByBrandSlice(salesByBrand map[string]saleByBrandItr) []SalesByBrand {
	result := make([]SalesByBrand, 0, len(salesByBrand))
	for brandName, salesData := range salesByBrand {
		result = append(result, SalesByBrand{
			BrandName:         brandName,
			MostSoldProduct:   getMostSoldProduct(salesData.MostSoldProduct),
			TotalQuantitySold: salesData.TotalQuantity,
			TotalRevenue:      salesData.TotalTransactionAmount,
			TotalProfit:       salesData.TotalProfit,
		})
	}
	return result
}

// Helper to get the most sold product
func getMostSoldProduct(productQuantities map[string]int) string {
	var maxProduct string
	var maxQuantity int
	for product, quantity := range productQuantities {
		if quantity > maxQuantity {
			maxProduct = product
			maxQuantity = quantity
		}
	}
	return maxProduct
}
