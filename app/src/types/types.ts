export type productType = {
    productId: number,
    productName: string,
    brandName: string,
    costPrice: number
    sellingPrice: number,
    category: string,
    ExpiryDate: string
}

export type salesType = {
    transactionId: number,
    productId: number,
    quantity: number,
    totalTransactionAmount: number,
    transactionDate: string
}