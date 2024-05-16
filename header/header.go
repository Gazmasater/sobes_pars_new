package header

import "github.com/tebeka/selenium"

func SetRequestHeaders(wd selenium.WebDriver) error {
	// Установить заголовок Accept
	if _, err := wd.ExecuteScript(`navigator.acceptHeader = "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,*/*;q=0.8";`, nil); err != nil {
		return err
	}

	// Установить заголовок User-Agent
	if _, err := wd.ExecuteScript(`navigator.userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.95 Safari/537.36";`, nil); err != nil {
		return err
	}

	// Установить заголовок Referer
	if _, err := wd.ExecuteScript(`document.referrer = "https://www.google.com/";`, nil); err != nil {
		return err
	}

	// Установить заголовок Accept-Encoding
	if _, err := wd.ExecuteScript(`navigator.acceptEncoding = "gzip, deflate, sdch";`, nil); err != nil {
		return err
	}

	// Установить заголовок Accept-Language
	if _, err := wd.ExecuteScript(`navigator.acceptLanguage = "ru-RU,ru;q=0.8,en-US,en;q=0.8";`, nil); err != nil {
		return err
	}

	return nil
}
