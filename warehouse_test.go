package main

import "testing"

func TestWarehouseAddProduct(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	err = warehouse.AddProduct(product)
	if err != nil {
		t.Fatal(err)
	}

	if len(warehouse.products) != 1 {
		t.Errorf("ожидали 1 товар, получили %d", len(warehouse.products))
	}
}

func TestWarehouseAddDuplicate(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(product); err != nil {
		t.Fatal(err)
	}

	err = warehouse.AddProduct(product)

	if err == nil {
		t.Error("ожидали ошибку при добавлении дубликата")
	}

	if len(warehouse.products) != 1 {
		t.Errorf("ожидали 1 товар, получили %d", len(warehouse.products))
	}
}

func TestWarehouseFindProduct(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Ноутбук", 50000, 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(product); err != nil {
		t.Fatal(err)
	}

	found, err := warehouse.FindProduct("Ноутбук")
	if err != nil {
		t.Fatal(err)
	}

	if found.Name != "Ноутбук" {
		t.Errorf("ожидали Ноутбук, получили %s", found.Name)
	}
}

func TestWarehouseFindMissingProduct(t *testing.T) {
	warehouse := NewWarehouse()

	_, err := warehouse.FindProduct("Планшет")

	if err == nil {
		t.Error("ожидали ошибку: товар не существует")
	}
}

func TestWarehouseSellProduct(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(product); err != nil {
		t.Fatal(err)
	}

	err = warehouse.SellProduct("Телефон", 4)
	if err != nil {
		t.Fatal(err)
	}

	found, err := warehouse.FindProduct("Телефон")
	if err != nil {
		t.Fatal(err)
	}

	if found.Quantity != 6 {
		t.Errorf("ожидали остаток 6, получили %d", found.Quantity)
	}
}

func TestWarehouseSellMissingProduct(t *testing.T) {
	warehouse := NewWarehouse()

	err := warehouse.SellProduct("Неизвестный", 2)

	if err == nil {
		t.Error("ожидали ошибку при продаже отсутствующего товара")
	}
}

func TestWarehouseRestockProduct(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Ноутбук", 50000, 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(product); err != nil {
		t.Fatal(err)
	}

	err = warehouse.RestockProduct("Ноутбук", 5)
	if err != nil {
		t.Fatal(err)
	}

	found, err := warehouse.FindProduct("Ноутбук")
	if err != nil {
		t.Fatal(err)
	}

	if found.Quantity != 10 {
		t.Errorf("ожидали остаток 10, получили %d", found.Quantity)
	}
}

func TestWarehouseRestockMissingProduct(t *testing.T) {
	warehouse := NewWarehouse()

	err := warehouse.RestockProduct("Неизвестный", 5)

	if err == nil {
		t.Error("ожидали ошибку при пополнении отсутствующего товара")
	}
}

func TestWarehouseDeleteProduct(t *testing.T) {
	warehouse := NewWarehouse()

	product, err := NewProduct("Планшет", 20000, 3)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(product); err != nil {
		t.Fatal(err)
	}

	err = warehouse.DeleteProduct("Планшет")
	if err != nil {
		t.Fatal(err)
	}

	if len(warehouse.products) != 0 {
		t.Errorf("ожидали пустой склад, получили %d товаров", len(warehouse.products))
	}
}

func TestWarehouseDeleteMissingProduct(t *testing.T) {
	warehouse := NewWarehouse()

	err := warehouse.DeleteProduct("Неизвестный")

	if err == nil {
		t.Error("ожидали ошибку удаления отсутствующего товара")
	}
}
