package services

type StockResponse struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result `json:"result"`
	Error  string   `json:"error"`
}

type Result struct {
	Meta Meta `json:"meta"`
}

type Meta struct {
	Currency           string  `json:"currency"`
	Symbol             string  `json:"symbol"`
	RegularMarketPrice float32 `json:"regularMarketPrice"`
}

type SearchResponse struct {
	Quotes []Quote `json:"quotes"`
}

type Quote struct {
	Exchange  string `json:"exchange"`
	Symbol    string `json:"symbol"`
	QuoteType string `json:"quoteType"`
	ShortName string `json:"shortname"`
}

type StockItem struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price"`
	Count  int64   `json:"count"`
}

func InsertStock(stockItem StockItem) (int64, error) {
	// db := repos.CreateConnection()
	// if err != nil {
	// 	return 0, err
	// }

	// return id, nil

	return 0, nil
}
