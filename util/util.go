package util

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/tebeka/selenium"
)

func FindFirstNumberIndex(text string) int {
	re := regexp.MustCompile(`\d+`)
	match := re.FindString(text)
	if match != "" {
		return strings.Index(text, match)
	}
	return -1
}

// Скриншот
func TakeScreenshot(wd selenium.WebDriver, filename string) error {
	// Сделать скриншот страницы
	screenshot, err := wd.Screenshot()
	if err != nil {
		return fmt.Errorf("ошибка при создании скриншота: %v", err)
	}

	// Сохранить скриншот в файл
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла: %v", err)
	}
	defer file.Close()

	_, err = file.Write(screenshot)
	if err != nil {
		return fmt.Errorf("ошибка при записи в файл: %v", err)
	}

	fmt.Printf("Скриншот сохранен в файле %s\n", filename)
	return nil
}
