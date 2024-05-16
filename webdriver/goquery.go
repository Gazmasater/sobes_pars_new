package webdriver

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
)

func CreateGoQueryDocumentFromPage(wd selenium.WebDriver) (*goquery.Document, error) {
	// Получаем HTML-код страницы
	html, err := wd.PageSource()
	if err != nil {
		fmt.Println("Ошибка при получении HTML-кода страницы:", err)
		return nil, err
	}

	// Создаем новый объект GoQuery для HTML-кода страницы
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		fmt.Println("Ошибка при создании GoQuery документа:", err)
		return nil, err
	}

	return doc, nil
}
