package product

type Product struct {
	ProductId    int    `json:"productId" bson:"productId"`
	ProductName  string `json:"productName" bson:"productName"`
	BrandName    string `json:"brandName" bson:"brandName"`
	CostPrice    int    `json:"costPrice" bson:"costPrice"`
	SellingPrice int    `json:"sellingPrice" bson:"sellingPrice"`
	Category     string `json:"category" bson:"category"`
	ExpiryDate   string `json:"expiryDate" bson:"expiryDate"`
}
