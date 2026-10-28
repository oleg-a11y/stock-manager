package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func (w *Warehouse) Save(filename string) error {
	data, err := json.MarshalIndent(w.products, "", "  ")
	if err != nil {
		return fmt.Errorf("ошибка преобразования данных: %w", err)
	}

	dir := filepath.Dir(filename)

	temp, err := os.CreateTemp(dir, ".stockmanager-*.tmp")
	if err != nil {
		return fmt.Errorf("ошибка создания файла: %w", err)
	}

	tempName := temp.Name()
	defer os.Remove(tempName)

	if _, err := temp.Write(data); err != nil {
		temp.Close()
		return fmt.Errorf("ошибка записи: %w", err)
	}

	if err := temp.Close(); err != nil {
		return fmt.Errorf("ошибка закрытия файла: %w", err)
	}

	if err := os.Rename(tempName, filename); err != nil {
		return fmt.Errorf("ошибка сохранения файла: %w", err)
	}

	return nil
}

func (w *Warehouse) Load(filename string) error {
	data, err := os.ReadFile(filename)

	if errors.Is(err, os.ErrNotExist) {
		return nil
	}

	if err != nil {
		return fmt.Errorf("ошибка чтения файла: %w", err)
	}

	var products []Product

	if err := json.Unmarshal(data, &products); err != nil {
		return fmt.Errorf("ошибка JSON: %w", err)
	}

	loaded := NewWarehouse()

	for _, product := range products {
		if err := loaded.AddProduct(product); err != nil {
			return fmt.Errorf("неверные данные в файле: %w", err)
		}
	}

	w.products = loaded.products
	return nil
}
