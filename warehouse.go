package main

import (
	"errors"
	"fmt"
)

type Warehouse struct {
	products []Product
}

func NewWarehouse() *Warehouse {
	return &Warehouse{
		products: make([]Product, 0),
	}
}

func (w *Warehouse) AddProduct(product Product) error {
	validated, err := NewProduct(
		product.Name, product.Price, product.Quantity,
	)
	if err != nil {
		return err
	}

	for _, existing := range w.products {
		if existing.Name == validated.Name {
			return fmt.Errorf("товар %q уже существует", validated.Name)
		}
	}

	w.products = append(w.products, validated)
	return nil
}

func (w *Warehouse) ListProducts() {
	if len(w.products) == 0 {
		fmt.Println("Склад пуст")
		return
	}

	for _, product := range w.products {
		product.Print()
	}
}

func (w *Warehouse) FindProduct(name string) (*Product, error) {
	for i := range w.products {
		if w.products[i].Name == name {
			return &w.products[i], nil
		}
	}

	return nil, fmt.Errorf("товар %q не найден", name)
}

func (w *Warehouse) SellProduct(name string, quantity int) error {
	product, err := w.FindProduct(name)
	if err != nil {
		return err
	}

	return product.Sell(quantity)
}

func (w *Warehouse) RestockProduct(name string, quantity int) error {
	product, err := w.FindProduct(name)
	if err != nil {
		return err
	}

	return product.Restock(quantity)
}

func (w *Warehouse) DeleteProduct(name string) error {
	for i, product := range w.products {
		if product.Name == name {
			w.products = append(
				w.products[:i],
				w.products[i+1:]...,
			)
			return nil
		}
	}

	return errors.New("товар не найден")
}
