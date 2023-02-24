package models

type GetSalesByMRes struct {
	SalesDate string `json:"sales_date"`
	Who       string `json:"who"`
	Name      string `json:"name"`
	SalesSum  string `json:"sales_sum"`
}

type GetSalesTop20 struct {
	Name     string `json:"name"`
	SalesSum string `json:"sales_sum"`
}

type GetSalesByMAccount struct {
	AccountNumber string `json:"account_number"`
	SalesDate     string `json:"sales_date"`
}

type GetSalesByProduct struct {
	ProductName string `json:"product_name"`
	SalesSum    string `json:"sales_sum"`
}
