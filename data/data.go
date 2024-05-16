package data

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

const (
	SeleniumPath = "/usr/local/bin/chromedriver"
	Port         = 4444
)

// WebDriverConfig содержит конфигурацию для запуска WebDriver
type WebDriverConfig struct {
	SeleniumPath  string              // Путь к ChromeDriver
	Port          int                 // Порт для WebDriver
	ChromeOptions chrome.Capabilities // Опции Chrome
	ProxyOptions  selenium.Proxy      // Опции прокси
}

var (
	CategoryLinks  = make(map[string]string) // Карта для хранения ссылок
	CategoryMap    = make(map[string]string) // Карта для соответствия дополнительных и основных категорий
	TextToImageURL = make(map[string]string) // Карта для соответствия основной категории и URL изображений
)
