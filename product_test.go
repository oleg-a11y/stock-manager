package main

import "testing"

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	if product.Name != "Телефон" {
		t.Errorf("ожидали Телефон, получили %s", product.Name)
	}

	if product.Price != 10000 {
		t.Errorf("ожидали цену 10000, получили %d", product.Price)
	}

	if product.Quantity != 10 {
		t.Errorf("ожидали остаток 10, получили %d", product.Quantity)
	}
}

func TestNewProductInvalidData(t *testing.T) {
	tests := []struct {
		name     string
		price    int
		quantity int
	}{
		{"", 100, 10},
		{"Телефон", 0, 10},
		{"Телефон", -100, 10},
		{"Телефон", 100, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProduct(tt.name, tt.price, tt.quantity)

			if err == nil {
				t.Error("ожидали ошибку, но получили nil")
			}
		})
	}
}

func TestProductSell(t *testing.T) {
	product, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	err = product.Sell(3)
	if err != nil {
		t.Fatal(err)
	}

	if product.Quantity != 7 {
		t.Errorf("ожидали 7 товаров, получили %d", product.Quantity)
	}
}

func TestProductSellInvalidQuantity(t *testing.T) {
	tests := []struct {
		name     string
		quantity int
	}{
		{"Больше остатка", 20},
		{"Ноль", 0},
		{"Отрицательное количество", -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := NewProduct("Телефон", 10000, 10)
			if err != nil {
				t.Fatal(err)
			}

			err = product.Sell(tt.quantity)

			if err == nil {
				t.Error("ожидали ошибку продажи")
			}

			if product.Quantity != 10 {
				t.Errorf("остаток изменился: получили %d", product.Quantity)
			}
		})
	}
}

func TestProductRestock(t *testing.T) {
	product, err := NewProduct("Ноутбук", 50000, 5)
	if err != nil {
		t.Fatal(err)
	}

	err = product.Restock(3)
	if err != nil {
		t.Fatal(err)
	}

	if product.Quantity != 8 {
		t.Errorf("ожидали 8 товаров, получили %d", product.Quantity)
	}
}

func TestProductRestockInvalidQuantity(t *testing.T) {
	tests := []int{0, -5}

	for _, quantity := range tests {
		product, err := NewProduct("Ноутбук", 50000, 5)
		if err != nil {
			t.Fatal(err)
		}

		err = product.Restock(quantity)

		if err == nil {
			t.Errorf("ожидали ошибку при пополнении на %d", quantity)
		}

		if product.Quantity != 5 {
			t.Errorf("остаток изменился: получили %d", product.Quantity)
		}
	}
}
