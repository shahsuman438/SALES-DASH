package reports

type SummeryReport struct {
	MostProfitableProduct  string `json:"mostProfitableProduct"`
	LeastProfitableProduct string `json:"leastProfitableProduct"`
	DateOfHighestSales     string `json:"dateOfHighestSales"`
	DateOfLeastSales       string `json:"dateOfLeastSales"`
}

type SalesByProduct struct {
	ProductName       string  `json:"productName"`
	BrandName         string  `json:"brandName"`
	Category          string  `json:"category"`
	TotalQuantitySold int     `json:"totalQuantitySold"`
	TotalRevenue      float64 `json:"totalRevenue"`
	TotalProfit       float64 `json:"totalProfit"`
}
type SalesByBrand struct {
	BrandName         string  `json:"brandName"`
	MostSoldProduct   string  `json:"mostSoldProduct"`
	TotalQuantitySold int     `json:"totalQuantitySold"`
	TotalRevenue      float64 `json:"totalRevenue"`
	TotalProfit       float64 `json:"totalProfit"`
}

type saleByProductItr struct {
	TotalQuantity          int     `json:"totalQuantity"`
	TotalTransactionAmount float64 `json:"totalTransactionAmount"`
}
type saleByBrandItr struct {
	MostSoldProduct        map[string]int `json:"mostSoldProduct"`
	TotalQuantity          int            `json:"totalQuantity"`
	TotalTransactionAmount float64        `json:"totalTransactionAmount"`
	TotalProfit            float64        `json:"totalProfit"`
}
