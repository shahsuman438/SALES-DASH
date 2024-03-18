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

export type salesByProductType = {
    productName: string
    brandName: string
    category: string
    totalQuantitySold: number
    totalRevenue: number
    totalProfit: number
}

export type salesByBrandType = {
    brandName: string
    mostSoldProduct: string
    totalQuantitySold: number
    totalRevenue: number
    totalProfit: number
}