package reports

type SummeryReport struct {
	MostProfitableProduct  string `json:"mostProfitableProduct"`
	LeastProfitableProduct string `json:"leastProfitableProduct"`
	DateOfHighestSales     string `json:"dateOfHighestSales"`
	DateOfLeastSales       string `json:"dateOfLeastSales"`
}

type SalesByProduct struct {
	Sn                int    `json:"sn"`
	ProductName       string `json:"productName"`
	BrandName         string `json:"brandName"`
	Category          string `json:"category"`
	TotalQuantitySold int    `json:"totalQuantitySold"`
	TotalRevenue      int    `json:"totalRevenue"`
	TotalProfit       int    `json:"totalProfit"`
}

type saleByProductItr struct {
	TotalQuantity          int `json:"totalQuantity"`
	TotalTransactionAmount int `json:"totalTransactionAmount"`
}
