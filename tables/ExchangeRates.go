package tables

type ExchangeRate struct {
	Id               int     `json:"id"`
	BaseCurrencyId   int     `json:"bcid"`
	TargetCurrencyId int     `json:"tcid"`
	Rate             float64 `json:"rate"`
}
