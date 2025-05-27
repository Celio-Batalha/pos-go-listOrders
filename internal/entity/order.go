package entity

import (
	"errors"
)

var (
	ErrInvalidID    = errors.New("ID is invalid")
	ErrInvalidPrice = errors.New("Price is invalid")
	ErrInvalidTax   = errors.New("Tax is invalid")
)

type Order struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return ErrInvalidID
	}
	if o.Price <= 0 {
		return ErrInvalidPrice
	}
	if o.Tax <= 0 {
		return ErrInvalidTax
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error {
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}
