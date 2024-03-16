package sales

type Sales struct {
	TransactionId          int    `json:"transactionId" bson:"transactionId"`
	ProductId              int    `json:"productId" bson:"productId"`
	Quantity               int    `json:"quantity" bson:"quantity"`
	TotalTransactionAmount int    `json:"totalTransactionAmount" bson:"totalTransactionAmount"`
	TransactionDate        string `json:"transactionDate" bson:"transactionDate"`
}
