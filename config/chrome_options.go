package config

import (
	"fmt"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func CreateChromeOptions(cfg *Config) chrome.Capabilities {
	options := chrome.Capabilities{
		Args: []string{
			//"--headless",
			//	"--private",
			fmt.Sprintf("--user-agent=%s", cfg.UserAgent), // Устанавливаем пользовательский агент
			"--window-size=1920,1080",                     // Устанавливаем размер окна браузера
			"--disable-gpu",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-automation",
			"--disable-extensions",
			"--disable-translate",

			"--disable-dev-shm-usage",
		},
	}

	return options
}

func CreateProxyOptions(cfg *Config) selenium.Proxy {
	proxyOpts := selenium.Proxy{
		Type: selenium.Manual,
		HTTP: cfg.ProxyHost,
	}
	return proxyOpts
}
