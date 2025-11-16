package struct 

type ExchangeRates struct {
	Id int `json:"id"`
	BaseCurrencyId int `json:"bcid"`
	TargetCurrencyId int `json:"tcid"`
	Rate double `json:"rate"`
}