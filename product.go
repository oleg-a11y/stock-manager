package main

import (
	"errors"
	"fmt"
	"strings"
)

type Product struct {
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}

func NewProduct(name string, price, quantity int) (Product, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return Product{}, errors.New("название товара не может быть пустым")
	}
	if price <= 0 {
		return Product{}, errors.New("цена должна быть больше нуля")
	}
	if quantity < 0 {
		return Product{}, errors.New("количество не может быть отрицательным")
	}

	return Product{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}, nil
}

func (p Product) Print() {
	fmt.Printf("Товар: %s | Цена: %d | Остаток: %d\n",
		p.Name, p.Price, p.Quantity)
}

func (p *Product) Sell(quantity int) error {
	if quantity <= 0 {
		return errors.New("количество для продажи должно быть больше нуля")
	}
	if quantity > p.Quantity {
		return fmt.Errorf("недостаточно товара %q: доступно %d",
			p.Name, p.Quantity)
	}

	p.Quantity -= quantity
	return nil
}

func (p *Product) Restock(quantity int) error {
	if quantity <= 0 {
		return errors.New("количество для пополнения должно быть больше нуля")
	}

	// Защита от переполнения int.
	maxInt := int(^uint(0) >> 1)
	if quantity > maxInt-p.Quantity {
		return errors.New("слишком большое количество товара")
	}

	p.Quantity += quantity
	return nil
}
