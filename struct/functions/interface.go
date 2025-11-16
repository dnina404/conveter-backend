package interface

import (
	"myapp/struct"
)

type Functions interface {
	ShowAll() []Currencies
	ShowOne(code string) *Currencies
	Post(Currencies) 
	GetExRates() []ExchangeRates
	GetRatesOfTwo() (*ExchangeRates, *ExchangeRates)
}