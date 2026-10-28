package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWarehouseSaveLoad(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "products.json")

	warehouse := NewWarehouse()

	phone, err := NewProduct("Телефон", 10000, 10)
	if err != nil {
		t.Fatal(err)
	}

	laptop, err := NewProduct("Ноутбук", 50000, 5)
	if err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(phone); err != nil {
		t.Fatal(err)
	}

	if err := warehouse.AddProduct(laptop); err != nil {
		t.Fatal(err)
	}

	if err := warehouse.Save(filename); err != nil {
		t.Fatal(err)
	}

	loadedWarehouse := NewWarehouse()

	if err := loadedWarehouse.Load(filename); err != nil {
		t.Fatal(err)
	}

	if len(loadedWarehouse.products) != 2 {
		t.Fatalf(
			"ожидали 2 товара, получили %d",
			len(loadedWarehouse.products),
		)
	}

	foundPhone, err := loadedWarehouse.FindProduct("Телефон")
	if err != nil {
		t.Fatal(err)
	}

	if foundPhone.Price != 10000 || foundPhone.Quantity != 10 {
		t.Errorf("данные телефона восстановлены неправильно: %+v", foundPhone)
	}

	foundLaptop, err := loadedWarehouse.FindProduct("Ноутбук")
	if err != nil {
		t.Fatal(err)
	}

	if foundLaptop.Price != 50000 || foundLaptop.Quantity != 5 {
		t.Errorf("данные ноутбука восстановлены неправильно: %+v", foundLaptop)
	}
}

func TestWarehouseSaveEmpty(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "products.json")

	warehouse := NewWarehouse()

	if err := warehouse.Save(filename); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}

	if string(data) != "[]" {
		t.Errorf("ожидали пустой JSON-массив, получили %s", data)
	}
}

func TestWarehouseLoadMissingFile(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "missing.json")

	warehouse := NewWarehouse()

	err := warehouse.Load(filename)

	if err != nil {
		t.Errorf("отсутствующий файл не должен вызывать ошибку: %v", err)
	}

	if len(warehouse.products) != 0 {
		t.Errorf("ожидали пустой склад, получили %d товаров", len(warehouse.products))
	}
}

func TestWarehouseLoadInvalidJSON(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "invalid.json")

	err := os.WriteFile(filename, []byte("{invalid json}"), 0600)
	if err != nil {
		t.Fatal(err)
	}

	warehouse := NewWarehouse()

	err = warehouse.Load(filename)

	if err == nil {
		t.Error("ожидали ошибку при загрузке неправильного JSON")
	}
}

func TestWarehouseLoadDuplicateProducts(t *testing.T) {
	filename := filepath.Join(t.TempDir(), "duplicates.json")

	data := []byte(`[
		{"name":"Телефон","price":10000,"quantity":10},
		{"name":"Телефон","price":20000,"quantity":5}
	]`)

	if err := os.WriteFile(filename, data, 0600); err != nil {
		t.Fatal(err)
	}

	warehouse := NewWarehouse()

	err := warehouse.Load(filename)

	if err == nil {
		t.Error("ожидали ошибку при загрузке дубликатов")
	}
}
