package util

import (
	"fmt"
	"io"

	"pars.com/config"

	"pars.com/data"
)

func PrintProductInfo(description, price, price_new string, writer io.Writer) {

	if price == price_new {

		fmt.Printf("Описание товара: %s\n", description)
		fmt.Fprintf(writer, "Описание товара: %s\n", description)

		fmt.Printf("Цена товара: %s\n", price)
		fmt.Fprintf(writer, "Цена товара : %s\n", price)
		fmt.Fprintf(writer, " %s\n", "")

		fmt.Println()

	} else {
		fmt.Printf("Описание товара: %s\n", description)
		fmt.Fprintf(writer, "Описание товара: %s\n", description)

		fmt.Printf("Цена товара старая: %s\n", price)
		fmt.Fprintf(writer, "Цена товара старая: %s\n", price)

		fmt.Printf("Цена товара новая: %s\n", price_new)
		fmt.Println()
		fmt.Fprintf(writer, "Цена товара новая: %s\n", price_new)
		fmt.Fprintf(writer, " %s\n", "")
	}

}

func PrintProductDetails(writer io.Writer, cfg *config.Config, href string) {
	fmt.Fprintf(writer, "Город доставки: %s\n", cfg.City)
	fmt.Printf("Город доставки: %s\n", cfg.City)

	fmt.Fprintf(writer, "Адрес доставки: %s\n", cfg.Street+","+cfg.HouseNumber)
	fmt.Printf("Адрес доставки: %s\n", cfg.Street+","+cfg.HouseNumber)

	a := data.CategoryMap[data.CategoryLinks[cfg.BaseURL]]
	fmt.Fprintf(writer, "Имя категории: %s\n", a)
	fmt.Printf("Имя категории: %s\n", data.CategoryMap[data.CategoryLinks[cfg.BaseURL]])

	fmt.Fprintf(writer, "Имя дополнительной категории: %s\n", data.CategoryLinks[cfg.BaseURL])
	fmt.Printf("Имя дополнительной категории: %s\n", data.CategoryLinks[cfg.BaseURL])

	fmt.Fprintf(writer, "Ссылка товара: %s\n", cfg.TwoURL+href)
	fmt.Printf("Ссылка товара: %s\n", cfg.TwoURL+href)

	fmt.Printf("Url изображения категории: %s\n", data.TextToImageURL[a])
	fmt.Fprintf(writer, "Url изображения категории: %s\n", data.TextToImageURL[a])
}
