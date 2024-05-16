package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-vgo/robotgo"

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

	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	fmt.Println("Ошибка при создании файла:", err)
	// 	return
	// }
	// defer file.Close()

	// writer := bufio.NewWriter(file)
	// defer writer.Flush()

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

	// Установка заголовков запроса
	if err := header.SetRequestHeaders(wd); err != nil {
		log.Fatalf("Failed to set request headers: %v", err)
	}

	// // Создаем клиент Redis
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "", // если у вас есть пароль
	// 	DB:       0,  // выберите базу данных
	// })
	// defer client.Close()

	// // Получаем сессионные данные из Redis
	// var ctx = context.Background()
	// sessionData, err := util.LoadSessionFromRedis(ctx, client, "session_key")
	// if err != nil {
	// 	if errors.Is(err, redis.Nil) {
	// 		fmt.Println("Session data not found in Redis. Skipping session loading.")
	// 	} else {
	// 		log.Printf(" to load session from Redis: %v", err)
	// 	}
	// } else {
	// 	println("Before applySessionToWebsite")

	// 	if err := util.ApplySessionToWebsite(wd, sessionData); err != nil {
	// 		log.Fatalf("Failed to apply session to website: %v", err)
	// 	}
	// }

	if err := wd.Get(cfg.BaseURL); err != nil {
		Logger.Error("Ошибка при загрузке URL:", zap.Error(err))
		return
	}

	// Сохраняем сессионные данные в Redis
	// sessionData = map[string]interface{}{
	// 	"user_id":   123,
	// 	"user_name": "john_doe",
	// }
	// jsonData, err := json.Marshal(sessionData)
	// if err != nil {
	// 	log.Fatalf("Failed to serialize session data: %v", err)
	// }

	// // Сохраняем данные сессии в Redis с ключом "session_key"
	// err = client.Set(ctx, "session_key", jsonData, 24*time.Hour).Err()
	// if err != nil {
	// 	log.Fatalf("Failed to save session data to Redis: %v", err)
	// }

	time.Sleep(10 * time.Second)

	elements, err := wd.FindElements(selenium.ByCSSSelector, "a")
	if err != nil {
		Logger.Error("Не найден элемент", zap.Error(err))
	}

	for _, elem := range elements {
		// Получаем значение атрибута "href"
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

	doc, err := goquery.NewDocument(cfg.BaseURL)
	if err != nil {
		log.Fatal(err)
	}
	println("ИЩЕМ!!!!")

	findElement(doc.Children(), "Нет, другой")

	// Ищем элементы, содержащие текст "Ваш город"

	// if err := wd.Get(cfg.BaseURL); err != nil {
	// 	fmt.Println("Ошибка загрузки страницы:", err)
	// 	return
	// }

	// _, err = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)
	// if err != nil {
	// 	log.Fatalf("Failed to scroll page: %v", err)
	// }

	// Получаем размеры экрана
	width, height := robotgo.GetScreenSize()
	println("Ширина экрана:", width)
	println("Высота экрана:", height)
	newX := width * 7 / 8 // 1/8 экрана от правого края
	newY := height / 2    // По середине по вертикали

	// Перемещаем курсор
	robotgo.Move(newX, newY)
	robotgo.Click()
	time.Sleep(15 * time.Second)

	filename := fmt.Sprintf("screenshot_%d.png", 1) // Создаем уникальное имя файла с помощью значения переменной i
	if err := util.TakeScreenshot(wd, filename); err != nil {
		log.Fatalf("Ошибка при создании скриншота после клика: %v", err)
	}
	println("1111111111111111", 1)

	time.Sleep(1000 * time.Second)

	// Ожидание загрузки элемента и клик по нему
	elements, err = wd.FindElements(selenium.ByCSSSelector, ".AddressConfirmBadge_buttons__Ou9hW > ._button--theme_secondary_10nio_51 ._text_7xv2z_4")
	if err != nil {
		fmt.Println("Ошибка при поиске элементов:", err)
		return
	}

	if err := util.TakeScreenshot(wd, "2.png"); err != nil {
		log.Fatalf("Ошибка при создании скриншота после клика: %v", err)
	}

	println("2222222222222222222")

	if len(elements) == 0 {
		fmt.Println("Элементы не найдены")
		return
	}

	// Нажатие на первый элемент, который соответствует селектору
	if err := elements[0].Click(); err != nil {
		fmt.Println("Ошибка клика по кнопке:", err)
		return
	}

	println("33333333333333333")

	time.Sleep(10 * time.Second)
	var input selenium.WebElement

	input, err = wd.FindElement(selenium.ByCSSSelector, "input._textInput_1frhv_1")
	if err != nil {
		log.Fatal(err)
	}

	println("444444444444444444")

	err = input.SendKeys(cfg.City)
	if err != nil {
		log.Fatal(err)
	}

	// Находим контейнер со списком предложенных вариантов Городов
	suggestContainer, err := wd.FindElement(selenium.ByCSSSelector, ".Suggest_root__KuclW")
	if err != nil {
		log.Fatal(err)
	}

	// Находим все элементы в списке предложенных вариантов Городов
	suggestedElements, err := suggestContainer.FindElements(selenium.ByCSSSelector, ".Suggest_suggestItem__hOaW9")
	if err != nil {
		log.Fatal(err)
	}

	// Выводим текст каждого элемента

	text := make([]string, 4)

	// Используйте цикл для обхода элементов и получения текста из каждого элемента
	var selectedElement selenium.WebElement

	// Проходим по списку предложенных элементов
	for i, elem := range suggestedElements {
		// Проверяем, чтобы не выйти за пределы массива text и не превысить максимальное количество элементов
		if i < 4 {
			// Получаем текст из элемента
			t, err := elem.Text()
			if err != nil {
				log.Fatal(err)
			}
			// Сохраняем текст в массив text
			text[i] = t
			// Проверяем, соответствует ли текст заданному городу
			if t == cfg.City {
				// Если условие истинно, выбираем этот элемент из предложенного списка
				selectedElement = elem
				// Нажимаем на выбранный элемент
				err := selectedElement.Click()
				if err != nil {
					log.Fatal(err)
				}
				break // Выходим из цикла, так как элемент найден и нажат
			}
		} else {
			// Если элементов больше не нужно, выходим из цикла
			break
		}
	}

	fmt.Println(text)

	time.Sleep(2 * time.Second) // Подождем немного, чтобы увидеть результат
	var inputElements []selenium.WebElement

	// Находим элементы input по CSS-селектору
	inputElements, err = wd.FindElements(selenium.ByCSSSelector, "input._textInput_1frhv_1")
	if err != nil {
		log.Fatal(err)
	}

	if len(inputElements) > 0 {
		// Используем первый элемент для ввода текста
		err := inputElements[1].SendKeys(cfg.Street + "," + cfg.HouseNumber)
		if err != nil {
			log.Fatal(err)
		}
	}

	time.Sleep(2 * time.Second) // Подождем немного, чтобы увидеть результат

	// Находим первый элемент в выпавшем списке
	firstSuggestionElement, err := wd.FindElement(selenium.ByCSSSelector, ".Suggest_root__KuclW:nth-child(1)")
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second) // Подождем немного, чтобы увидеть результат

	// Выполняем действие с найденным элементом, например, нажатие на него
	err = firstSuggestionElement.Click()
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(2 * time.Second) // Подождем немного, чтобы увидеть результат

	// Находим элемент по CSS-селектору
	buttonElement, err := wd.FindElement(selenium.ByCSSSelector, "span[data-tid='text']._text_7xv2z_4._text--type_p1SemiBold_7xv2z_109._text_10nio_43")
	if err != nil {
		log.Fatal(err)
	}

	// Нажимаем на элемент
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

	// Находим все основные категории и соответствующие им дополнительные категории
	Doc.Find(".CatalogTreeSectionCard_name__DKRiD").Each(func(i int, s *goquery.Selection) {
		mainCategory := s.Text()

		// Находим родительский элемент, содержащий как основную, так и дополнительную категории
		parent := s.Parent()

		// Получаем текст следующего элемента после родительского (дополнительная категория)
		additionalCategories := parent.Next().Text()

		// Заменяем "Пасха" на "пасха"
		additionalCategories = strings.Replace(additionalCategories, "Пасха", "пасха", -1)

		// Разделяем строку с дополнительными категориями по регулярному выражению
		re := regexp.MustCompile(`[А-Я][^А-Я]*`)
		additionalCategoriesArr := re.FindAllString(additionalCategories, -1)

		// Добавляем каждую категорию в карту
		for _, category := range additionalCategoriesArr {
			category = strings.TrimSpace(category)
			additionalCategories = strings.TrimSuffix(additionalCategories, ",")
			if category != "" {
				data.CategoryMap[category] = mainCategory
			}
		}
	})

	// Создаем регулярное выражение для поиска URL изображения в атрибуте style
	re := regexp.MustCompile(`url\("(.*?)"\)`)

	// Объявляем массивы для хранения URL изображений и текста
	var imageURLs []string
	var texts []string

	// Обработка элементов с классом .CatalogTreeSectionCard_image__uobnI
	Doc.Find(".CatalogTreeSectionCard_image__uobnI").Each(func(i int, s *goquery.Selection) {
		// Получаем значение атрибута style
		style, exists := s.Attr("style")

		if exists {
			// Ищем URL изображения в атрибуте style с помощью регулярного выражения
			matches := re.FindStringSubmatch(style)
			if len(matches) >= 2 {
				imageURL := matches[1]
				// Добавляем URL изображения в массив imageURLs
				imageURLs = append(imageURLs, imageURL)
			}
		}
	})

	// Обработка элементов с классом .CatalogTreeSectionCard_name__DKRiD
	Doc.Find(".CatalogTreeSectionCard_name__DKRiD").Each(func(i int, s *goquery.Selection) {
		// Получаем текстовое содержимое элемента
		text := s.Text()

		// Проверяем, содержит ли текст слово "Пасха"
		text = strings.Replace(text, "Пасха", "пасха", -1)

		// Добавляем текст в массив texts
		texts = append(texts, text)
	})

	// Заполняем карту значениями из массивов imageURLs и texts
	for i, text := range texts {
		// Проверяем, есть ли URL изображения для данного текста
		if i < len(imageURLs) {
			data.TextToImageURL[text] = imageURLs[i]

		}
	}

	slideElements := Doc.Find(".ProductCard_root__OBGd_")
	var price string // Объявляем переменную здесь

	// Проходимся по каждому найденному элементу и его дочерним элементам
	slideElements.Each(func(i int, s *goquery.Selection) {
		text := s.Text()
		re := regexp.MustCompile(`\d+`)

		matches := re.FindAllString(text, -1)
		if len(matches) > 0 {
			numericPrice := matches[len(matches)-1] // Последний элемент массива - это наша цена
			switch len(numericPrice) {
			case 4:
				//fmt.Printf("Цена_старая: %s\n", numericPrice[:2])
				price = numericPrice[:2]
			case 5:
				//fmt.Printf("Цена_старая: %s\n", numericPrice[:3])
				price = numericPrice[:3]

			case 6:
				//fmt.Printf("Цена_старая: %s\n", numericPrice[:3])
				price = numericPrice[:3]

			default:
				//fmt.Printf("Цена: %s\n", numericPrice)
				price = numericPrice

			}
		} else {
			fmt.Println("Цена не найдена")
		}

		// Получаем родительский элемент текущего элемента
		parent := s.Parent()

		// Проверяем, является ли родительский элемент ссылкой
		if parent.Is("a") {
			// Получаем атрибут href родительского элемента, если он есть
			href, exists := parent.Attr("href")
			if exists {
				fmt.Println("Город доставки:", cfg.City)
				fmt.Println("Адрес доставки:", cfg.Street+","+cfg.HouseNumber)
				fmt.Println("Имя категории:", data.CategoryMap[data.CategoryLinks[cfg.BaseURL]])
				fmt.Println("Имя дополнительной категории:", data.CategoryLinks[cfg.BaseURL])
				fmt.Println("Ссылка товара:", cfg.TwoURL+href)
			}

		}

		// Выводим другие элементы внутри текущего элемента
		s.Children().Each(func(j int, child *goquery.Selection) {
			// Находим индекс первого числа в тексте описания
			childtext := child.Text()

			firstNumberIndex := util.FindFirstNumberIndex(childtext)

			// Если найдено число в тексте описания
			if firstNumberIndex >= 0 {
				// Обрезаем текст описания до первого числа
				description := child.Text()[:firstNumberIndex]
				if len(description) > 3 {
					fmt.Printf("Описание товара: %s\n", description)
					fmt.Printf("Цена товара: %s\n", price)

				}
			}

			// Получаем и выводим все изображения внутри текущего дочернего элемента
			child.Find("img").Each(func(k int, img *goquery.Selection) {
				// Получаем и выводим URL изображения
				src, exists := img.Attr("src")
				if exists {
					fmt.Println("URL изображения:", src)
				}
			})
		})
		fmt.Println()
	})

	println()

	// //	Пауза для просмотра результатов
	time.Sleep(2 * time.Second)

}

func findElement(selection *goquery.Selection, searchText string) {
	selection.Each(func(i int, s *goquery.Selection) {
		// Проверяем текущий элемент
		s.Each(func(_ int, attr *goquery.Selection) {
			if attr.Nodes != nil {
				for _, node := range attr.Nodes {
					for _, attr := range node.Attr {
						if attr.Key == "data-tid" {
							fmt.Println("data-tid:", attr.Val)
						}
					}
				}
			}
		})

		// Рекурсивно вызываем функцию для поиска в дочерних элементах текущего элемента
		findElement(s.Children(), searchText)
	})
}
