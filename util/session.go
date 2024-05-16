package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/tebeka/selenium"
)

func LoadSessionFromRedis(ctx context.Context, client *redis.Client, key string) (map[string]interface{}, error) {
	// Получаем данные сессии из Redis
	jsonData, err := client.Get(ctx, key).Bytes()
	if err != nil {
		return nil, fmt.Errorf("failed to get session data from Redis: %v", err)
	}

	// Декодируем данные сессии из JSON
	var sessionData map[string]interface{}
	err = json.Unmarshal(jsonData, &sessionData)
	if err != nil {
		return nil, fmt.Errorf("failed to deserialize session data: %v", err)
	}

	// Выводим содержимое мапы
	fmt.Println("Session Data:")
	for key, value := range sessionData {
		fmt.Printf("%s: %v\n", key, value)
	}

	return sessionData, nil
}

func ApplySessionToWebsite(wd selenium.WebDriver, sessionData map[string]interface{}) error {
	userID, ok := sessionData["user_id"].(float64)
	if !ok {
		return errors.New("failed to get user ID from session data")
	}
	userName, ok := sessionData["user_name"].(string)
	if !ok {
		return errors.New("failed to get user name from session data")
	}

	// Применяем сессию на веб-сайте, например, передавая данные в HTTP-заголовках
	cookieUserID := &selenium.Cookie{
		Name:   "user_id",
		Value:  fmt.Sprintf("%d", int(userID)),
		Domain: "",  // Пустая строка для применения к текущему домену
		Path:   "/", // Куки будут доступны на всех страницах
	}

	err := wd.AddCookie(cookieUserID)
	if err != nil {
		return fmt.Errorf("failed to add user ID cookie: %v", err)
	}

	cookieUserName := &selenium.Cookie{
		Name:   "user_name",
		Value:  userName,
		Domain: "",
		Path:   "/", // Куки будут доступны на всех страницах
	}

	err = wd.AddCookie(cookieUserName)
	if err != nil {
		return fmt.Errorf("failed to add user name cookie: %v", err)
	}

	return nil
}
