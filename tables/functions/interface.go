package interfaces

import (
	"Conveter/tables"
	"errors"
)

type Currencies struct {
	currcs []tables.Currency
}

type Functions interface {
	ShowAll() []tables.Currency
	ShowOne(code string) (*tables.Currency, error)
	Change(id1, id2 int) (float32, error)
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

func (c Currencies) Change(id1, id2 int) (float32, error) {
	var rate1 float32 = 0
	var rate2 float32 = 0
	for i := range len(c.currcs) {
		if i == id1 {
			rate1 = c.currcs[i].ToDollar
		}
		if i == id2 {
			rate2 = c.currcs[i].ToDollar
		}
	}
	if rate1 == 0 || rate2 == 0 {
		err := errors.New("cant find currency with that id")
		return 0, err
	}
	return rate1 / rate2, nil

}
