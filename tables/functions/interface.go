package interfaces

import (
	"errors"
	"myapp/tables"
)

type Currencies struct {
	currcs []tables.Currency
}
type ExchangeRates struct {
	rattes []tables.ExchangeRate
}

type Functions interface {
	ShowAll() []tables.Currency
	ShowOne(code string) (*tables.Currency, error)
	Post(*tables.Currency)
	GetExRates() []ExchangeRates
	GetRatesOfTwo() (ExchangeRates, ExchangeRates)
}

func (c Currencies) ShowAll() []tables.Currency {
	return c.currcs
}

func (c Currencies) ShowOne(code string) (*tables.Currency, error) {
	for i := range len(c.currcs) {
		if code == c.currcs[i].Code {
			return &c.currcs[i], nil
		}
	}
	err := errors.New("currency with this code is not exist")
	return nil, err

}
