package reports

type SummeryReport struct {
	MostProfitableProduct  string `json:"mostProfitableProduct"`
	LeastProfitableProduct string `json:"leastProfitableProduct"`
	DateOfHighestSales     string `json:"dateOfHighestSales"`
	DateOfLeastSales       string `json:"dateOfLeastSales"`
}
