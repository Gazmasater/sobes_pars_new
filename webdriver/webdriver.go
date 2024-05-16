package webdriver

import (
	"fmt"

	"github.com/tebeka/selenium"
	"pars.com/data"
)

// StartWebDriver запускает WebDriver с заданной конфигурацией
func StartWebDriver(config data.WebDriverConfig) (selenium.WebDriver, error) {
	// Настройка сервиса ChromeDriver
	opts := []selenium.ServiceOption{}
	_, err := selenium.NewChromeDriverService(config.SeleniumPath, config.Port, opts...)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запуске Chrome WebDriver: %s", err)
	}

	// Создание удаленного драйвера с настроенными параметрами прокси
	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(config.ChromeOptions)

	wd, err := selenium.NewRemote(
		caps, // Передаем capabilities
		fmt.Sprintf("http://localhost:%d/wd/hub", config.Port),
	)

	if err != nil {
		return nil, fmt.Errorf("ошибка при создании удаленного драйвера: %s", err)
	}

	return wd, nil
}
