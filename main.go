package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// Открываем HTML-файл
	file, err := os.Open("page_source.html")
	if err != nil {
		log.Fatalf("Error opening HTML file: %s\n", err)
	}
	defer file.Close()

	// Загружаем HTML-документ с использованием goquery
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalf("Error loading HTML document: %s\n", err)
	}

	// Ищем все элементы с классом "ambrands-label"
	doc.Find("span.ambrands-label").Each(func(i int, s *goquery.Selection) {
		// Извлекаем и выводим текст из каждого найденного элемента
		text := s.Text()
		fmt.Printf("Brand %d: %s\n", i+1, text)
	})
}
