package models

type Currency struct {
	id       int    `json:"id"`
	Code     string `json:"code"`
	FullName string `json:"fullname"`
	Sign     string `json:"sign"`
}

type ExchangeRate struct {
	Id               int     `json:"id"`
	BaseCurrencyId   int     `json:"bcid"`
	TargetCurrencyId int     `json:"tcid"`
	Rate             float64 `json:"rate"`
}
