Тем не менее, существует обходной путь: можно сначала открыть пустую страницу (about:blank), выполнить необходимый JavaScript для изменения window.navigator.webdriver, а затем уже загрузить целевую страницу.

Вот пример реализации этого подхода:

go
Копировать код
package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func StartWebDriver() (selenium.WebDriver, error) {
	const (
		seleniumPath = "/path/to/selenium-server.jar" // Укажите путь к вашему Selenium серверу
		port         = 8080
	)

	opts := []selenium.ServiceOption{}
	_, err := selenium.NewChromeDriverService(seleniumPath, port, opts...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запуске Chrome WebDriver: %s", err)
	}

	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Path: "",
		Args: []string{
			"--headless",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании удаленного драйвера: %s", err)
	}

	return wd, nil
}

func main() {
	wd, err := StartWebDriver()
	if err != nil {
		log.Fatalf("ошибка при запуске WebDriver: %s", err)
	}
	defer wd.Quit()

	// Откройте пустую страницу
	err = wd.Get("about:blank")
	if err != nil {
		log.Fatalf("ошибка при открытии пустой страницы: %s", err)
	}

	// Выполните JavaScript для изменения атрибута перед загрузкой целевой страницы
	jsCode := `Object.defineProperty(navigator, 'webdriver', {get: () => false});`
	_, err = wd.ExecuteScript(jsCode, nil)
	if err != nil {
		log.Fatalf("ошибка при выполнении JavaScript: %s", err)
	}

	// Теперь откройте целевую страницу
	err = wd.Get("http://example.com") // Замените URL на необходимый
	if err != nil {
		log.Fatalf("ошибка при открытии целевой страницы: %s", err)
	}

	// Выполните дополнительные действия с браузером
}
Объяснение:
Открытие пустой страницы:

Вначале открывается пустая страница с помощью wd.Get("about:blank"), чтобы предотвратить загрузку целевой страницы до выполнения изменений.
Изменение атрибута:

После этого выполняется JavaScript для изменения значения атрибута window.navigator.webdriver.
Загрузка целевой страницы:

Затем, после изменения атрибута, загружается нужная веб-страница.
Заключение:
Этот подход гарантирует, что атрибут будет изменён до того, как целевая страница будет загружена и сможет обнаружить автоматизацию. Это снижает вероятность того, что веб-сайт сможет определить использование WebDriver до внесения изменений.
