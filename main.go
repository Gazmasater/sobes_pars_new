package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"go.uber.org/zap"
	"pars.com/config"
	"pars.com/data"
	"pars.com/header"
	"pars.com/util"
	"pars.com/webdriver"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Ошибка при загрузке конфигурации:", err)
		return
	}

	Log_init()

	configg := data.WebDriverConfig{
		SeleniumPath:  data.SeleniumPath,
		Port:          data.Port,
		ChromeOptions: config.CreateChromeOptions(cfg),
		ProxyOptions:  config.CreateProxyOptions(cfg),
	}
	wd, err := webdriver.StartWebDriver(configg)
	if err != nil {
		Logger.Error("Ошибка при запуске WebDriver:", zap.Error(err))
		return
	}
	defer wd.Quit()

	if err := header.SetRequestHeaders(wd); err != nil {
		log.Fatalf("Failed to set request headers: %v", err)
	}

	if err := wd.Get(cfg.BaseURL); err != nil {
		Logger.Error("Ошибка при загрузке URL:", zap.Error(err))
		return
	}

	elements, err := wd.FindElements(selenium.ByCSSSelector, "a")
	if err != nil {
		Logger.Error("Не найден элемент", zap.Error(err))
	}

	for _, elem := range elements {
		link, err := elem.GetAttribute("href")
		if err != nil {
			fmt.Printf("Ошибка при получении атрибута href: %v\n", err)
			continue
		}

		if strings.Contains(link, "category") {
			categoryName, err := elem.Text()
			if categoryName == "Скоро Пасха" {
				categoryName = "Скоро пасха"
			}

			if err != nil {
				fmt.Printf("Ошибка при получении текста элемента: %v\n", err)
				continue
			}

			data.CategoryLinks[link] = categoryName
		}
	}

	elements, err = wd.FindElements(selenium.ByCSSSelector, ".AddressConfirmBadge_buttons__Ou9hW > ._button--theme_secondary_10nio_51 ._text_7xv2z_4")
	if err != nil {
		fmt.Println("Ошибка при поиске элементов:", err)
		return
	}

	if len(elements) == 0 {
		fmt.Println("Элементы не найдены")
		return
	}

	if err := elements[0].Click(); err != nil {
		fmt.Println("Ошибка клика по кнопке:", err)
		return
	}

	time.Sleep(3 * time.Second)
	var input selenium.WebElement

	input, err = wd.FindElement(selenium.ByCSSSelector, "input._textInput_1frhv_1")
	if err != nil {
		log.Fatal(err)
	}

	err = input.SendKeys(cfg.City)
	if err != nil {
		log.Fatal(err)
	}

	suggestContainer, err := wd.FindElement(selenium.ByCSSSelector, ".Suggest_root__KuclW")
	if err != nil {
		log.Fatal(err)
	}

	suggestedElements, err := suggestContainer.FindElements(selenium.ByCSSSelector, ".Suggest_suggestItem__hOaW9")
	if err != nil {
		log.Fatal(err)
	}

	text := make([]string, 4)

	var selectedElement selenium.WebElement

	for i, elem := range suggestedElements {
		if i < 4 {
			t, err := elem.Text()
			if err != nil {
				log.Fatal(err)
			}
			text[i] = t
			if t == cfg.City {
				selectedElement = elem
				err := selectedElement.Click()
				if err != nil {
					log.Fatal(err)
				}
				break
			}
		} else {
			break
		}
	}

	fmt.Println(text)

	time.Sleep(2 * time.Second)
	var inputElements []selenium.WebElement

	inputElements, err = wd.FindElements(selenium.ByCSSSelector, "input._textInput_1frhv_1")
	if err != nil {
		log.Fatal(err)
	}

	if len(inputElements) > 0 {
		err := inputElements[1].SendKeys(cfg.Street + "," + cfg.HouseNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(2 * time.Second)

	firstSuggestionElement, err := wd.FindElement(selenium.ByCSSSelector, ".Suggest_root__KuclW:nth-child(1)")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	err = firstSuggestionElement.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	buttonElement, err := wd.FindElement(selenium.ByCSSSelector, "span[data-tid='text']._text_7xv2z_4._text--type_p1SemiBold_7xv2z_109._text_10nio_43")
	if err != nil {
		log.Fatal(err)
	}

	err = buttonElement.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second)

	if err := wd.Get(cfg.BaseURL); err != nil {
		Logger.Error("Ошибка при загрузке URL:", zap.Error(err))
		return
	}

	Doc, err := webdriver.CreateGoQueryDocumentFromPage(wd)

	if err != nil {
		Logger.Error("goguery документ не создан", zap.Error(err))
	}

	Doc.Find(".CatalogTreeSectionCard_name__DKRiD").Each(func(i int, s *goquery.Selection) {
		mainCategory := s.Text()
		parent := s.Parent()
		additionalCategories := parent.Next().Text()
		additionalCategories = strings.Replace(additionalCategories, "Пасха", "пасха", -1)
		re := regexp.MustCompile(`[А-Я][^А-Я]*`)
		additionalCategoriesArr := re.FindAllString(additionalCategories, -1)

		for _, category := range additionalCategoriesArr {
			category = strings.TrimSpace(category)
			additionalCategories = strings.TrimSuffix(additionalCategories, ",")
			if category != "" {
				data.CategoryMap[category] = mainCategory
			}
		}
	})

	file_categ := fmt.Sprintf("%s.txt", data.CategoryMap[data.CategoryLinks[cfg.BaseURL]])

	file, err := os.Create(file_categ)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	re := regexp.MustCompile(`url\("(.*?)"\)`)

	var imageURLs []string
	var texts []string

	Doc.Find(".CatalogTreeSectionCard_image__uobnI").Each(func(i int, s *goquery.Selection) {
		style, exists := s.Attr("style")

		if exists {
			matches := re.FindStringSubmatch(style)
			if len(matches) >= 2 {
				imageURL := matches[1]
				imageURLs = append(imageURLs, imageURL)
			}
		}
	})

	Doc.Find(".CatalogTreeSectionCard_name__DKRiD").Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		text = strings.Replace(text, "Пасха", "пасха", -1)
		texts = append(texts, text)
	})

	for i, text := range texts {
		if i < len(imageURLs) {
			data.TextToImageURL[text] = imageURLs[i]
		}
	}

	slideElements := Doc.Find(".ProductCard_root__OBGd_")
	var price string

	slideElements.Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		re := regexp.MustCompile(`\d+`)
		matches := re.FindAllString(text, -1)

		var numericPrice string
		if len(matches) > 0 {
			numericPrice = matches[len(matches)-1]
			switch len(numericPrice) {
			case 4:
				price = numericPrice[:2]
			case 5:
				price = numericPrice[:3]
			case 6:
				price = numericPrice[:3]
			default:
				price = numericPrice
			}
		} else {
			fmt.Println("Цена не найдена")
			fmt.Fprintln(writer, "Цена не найдена")
			return
		}

		parent := s.Parent()
		if !parent.Is("a") {
			fmt.Printf("Родительский элемент не является ссылкой для товара %d\n", i)
			fmt.Fprintf(writer, "Родительский элемент не является ссылкой для товара %d\n", i)
			return
		}

		href, exists := parent.Attr("href")
		if !exists {
			fmt.Printf("Атрибут href не найден для товара %d\n", i)
			fmt.Fprintf(writer, "Атрибут href не найден для товара %d\n", i)
			return
		}

		fmt.Fprintf(writer, "Город доставки: %s\n", cfg.City)
		fmt.Printf("Город доставки: %s\n", cfg.City)

		fmt.Fprintf(writer, "Адрес доставки: %s\n", cfg.Street+","+cfg.HouseNumber)
		fmt.Printf("Адрес доставки: %s\n", cfg.Street+","+cfg.HouseNumber)

		fmt.Fprintf(writer, "Имя категории: %s\n", data.CategoryMap[data.CategoryLinks[cfg.BaseURL]])
		fmt.Printf("Имя категории: %s\n", data.CategoryMap[data.CategoryLinks[cfg.BaseURL]])

		fmt.Fprintf(writer, "Имя дополнительной категории: %s\n", data.CategoryLinks[cfg.BaseURL])
		fmt.Printf("Имя дополнительной категории: %s\n", data.CategoryLinks[cfg.BaseURL])

		fmt.Fprintf(writer, "Ссылка товара: %s\n", cfg.TwoURL+href)
		fmt.Printf("Ссылка товара: %s\n", href)

		s.Children().Each(func(j int, child *goquery.Selection) {
			childtext := child.Text()
			firstNumberIndex := util.FindFirstNumberIndex(childtext)
			if firstNumberIndex >= 0 {
				description := child.Text()[:firstNumberIndex]
				if len(description) > 3 {
					fmt.Printf("Описание товара: %s\n", description)
					fmt.Fprintf(writer, "Описание товара: %s\n", description)

					fmt.Printf("Цена товара: %s\n", price)
					fmt.Println()
					fmt.Fprintf(writer, "Цена товара: %s\n", price)
					fmt.Fprintf(writer, " %s\n", "")

				}
			}

			child.Find("img").Each(func(k int, img *goquery.Selection) {
				src, exists := img.Attr("src")
				if exists {
					fmt.Fprintf(writer, "URL изображения: %s\n", src)
					fmt.Println("URL изображения:", src)

				}
			})
		})
	})
	writer.Flush()

	time.Sleep(2000 * time.Second)
}
