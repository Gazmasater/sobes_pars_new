package util

import (
	"fmt"
	"io"
	"strings"
	"time"

	"pars.com/config"

	"pars.com/data"
)

func PrintProductInfo(description, price, pricenew string, writer io.Writer) {

	if price == pricenew {

		//.Printf("Описание товара: %s\n", description)
		fmt.Fprintf(writer, "Описание товара: %s\n", description)

		//fmt.Printf("Цена товара: %s\n", price)
		fmt.Fprintf(writer, "Цена товара : %s\n", price)
		fmt.Fprintf(writer, " %s\n", "")

		//	fmt.Println()

	} else {
		//	fmt.Printf("Описание товара: %s\n", description)
		fmt.Fprintf(writer, "Описание товара: %s\n", description)

		//	fmt.Printf("Цена товара старая: %s\n", price)
		fmt.Fprintf(writer, "Цена товара старая: %s\n", price)

		//	fmt.Printf("Цена товара новая: %s\n", price_new)
		//	fmt.Println()
		fmt.Fprintf(writer, "Цена товара новая: %s\n", pricenew)
		fmt.Fprintf(writer, " %s\n", "")
	}

}

func PrintProductDetails(writer io.Writer, cfg *config.Config, href string) {
	fmt.Fprintf(writer, "Город доставки: %s\n", cfg.City)
	fmt.Fprintf(writer, "Адрес доставки: %s\n", cfg.Street+","+cfg.HouseNumber)

	a1 := strings.TrimSpace(data.CategoryLinks[cfg.BaseURL])
	fmt.Printf("a1: '%s'\n", a1) // Отладочный вывод значения a1

	// Проверка наличия ключа a1 в мапе data.CategoryMap и вывод содержимого
	a, exists := data.CategoryMap[a1]
	if exists {
		fmt.Printf("CategoryMap[%s]:\n", a1)
		for key, value := range a {
			fmt.Printf("  %d: %c\n", key, value) // Change %s to %c for rune value
		}
	} else {
		fmt.Printf("CategoryMap[%s] не существует или пуст\n", a1)
	}

	// Дополнительный вывод для отладки
	fmt.Printf("Содержимое data.CategoryMap:\n")
	for key, value := range data.CategoryMap {
		fmt.Printf("  %s: %v\n", key, value)
	}

	time.Sleep(100 * time.Second)

	fmt.Fprintf(writer, "Имя категории: %s\n", a)
	fmt.Fprintf(writer, "Имя дополнительной категории: %s\n", a1)
	fmt.Fprintf(writer, "Ссылка товара: %s\n", cfg.TwoURL+href)
	fmt.Fprintf(writer, "Url изображения категории: %s\n", data.TextToImageURL[a])
}
