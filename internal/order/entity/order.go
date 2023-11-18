package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func (o Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid ID")
	}

	if o.Price == 0 {
		return errors.New("invalid Price")
	}

	if o.Tax == 0 {
		return errors.New("invalid Tax")
	}
	return nil
}
