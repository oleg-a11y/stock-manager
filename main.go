package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const dataFile = "products.json"

func main() {
	warehouse := NewWarehouse()

	if err := warehouse.Load(dataFile); err != nil {
		fmt.Println("Ошибка загрузки:", err)
		return
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n===== STOCK MANAGER =====")
		fmt.Println("1. Показать все товары")
		fmt.Println("2. Добавить товар")
		fmt.Println("3. Найти товар")
		fmt.Println("4. Продать товар")
		fmt.Println("5. Пополнить запасы")
		fmt.Println("6. Удалить товар")
		fmt.Println("0. Выход")
		fmt.Print("Выберите действие: ")

		if !scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Println("Ошибка ввода:", err)
			}
			return
		}

		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			warehouse.ListProducts()

		case "2":
			name, ok := readText(scanner, "Название: ")
			if !ok {
				return
			}

			price, ok := readInt(scanner, "Цена: ")
			if !ok {
				continue
			}

			quantity, ok := readInt(scanner, "Количество: ")
			if !ok {
				continue
			}

			product, err := NewProduct(name, price, quantity)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			err = warehouse.AddProduct(product)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			if !saveChanges(warehouse) {
				return
			}
			fmt.Println("Товар добавлен")

		case "3":
			name, ok := readText(scanner, "Название товара: ")
			if !ok {
				return
			}

			product, err := warehouse.FindProduct(name)
			if err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			product.Print()

		case "4":
			name, ok := readText(scanner, "Название товара: ")
			if !ok {
				return
			}

			quantity, ok := readInt(scanner, "Сколько продать: ")
			if !ok {
				continue
			}

			if err := warehouse.SellProduct(name, quantity); err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			if !saveChanges(warehouse) {
				return
			}
			fmt.Println("Продажа выполнена")

		case "5":
			name, ok := readText(scanner, "Название товара: ")
			if !ok {
				return
			}

			quantity, ok := readInt(scanner, "Сколько добавить: ")
			if !ok {
				continue
			}

			if err := warehouse.RestockProduct(name, quantity); err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			if !saveChanges(warehouse) {
				return
			}
			fmt.Println("Запасы пополнены")

		case "6":
			name, ok := readText(scanner, "Название товара: ")
			if !ok {
				return
			}

			if err := warehouse.DeleteProduct(name); err != nil {
				fmt.Println("Ошибка:", err)
				continue
			}

			if !saveChanges(warehouse) {
				return
			}
			fmt.Println("Товар удалён")

		case "0":
			fmt.Println("До свидания!")
			return

		default:
			fmt.Println("Неизвестная команда")
		}
	}
}

func readText(scanner *bufio.Scanner, prompt string) (string, bool) {
	fmt.Print(prompt)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Println("Ошибка ввода:", err)
		}
		return "", false
	}

	return strings.TrimSpace(scanner.Text()), true
}

func readInt(scanner *bufio.Scanner, prompt string) (int, bool) {
	text, ok := readText(scanner, prompt)
	if !ok {
		return 0, false
	}

	value, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println("Ошибка: введите целое число")
		return 0, false
	}

	return value, true
}

func saveChanges(warehouse *Warehouse) bool {
	if err := warehouse.Save(dataFile); err != nil {
		fmt.Println("Критическая ошибка сохранения:", err)
		fmt.Println("Программа остановлена, чтобы не продолжать с несохранёнными данными.")
		return false
	}

	return true
}
