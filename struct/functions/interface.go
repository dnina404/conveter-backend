package interface

import (
	"myapp/struct"
	"error"
)

type Functions interface {
	ShowAll() Currencies.currcs
	ShowOne(code string) (Currencies, error)
	Post(Currency) 
	GetExRates() []ExchangeRates
	GetRatesOfTwo() (ExchangeRates, ExchangeRates)
}

type Currencies struct {
	currcs []struct.Currency
}
type ExchangeRates struct {
	rattes []struct.ExchangeRate
}

func (c Currencies) ShowAll() Currencies.currcs {
	return &c.currcs
}

func (c Currencies) ShowOne(code string) (Currencies, error) {
	for i := range(len(c.currcs)) {
		if code == c.currcs[i].code {
			return c.currcs[i], nil 
		}
	}
	err := error.NewError("non existent currency")
	return nil, err

}
