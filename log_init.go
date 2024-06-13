package main

import "go.uber.org/zap"

var Logger *zap.Logger

func Loginit() {
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic(err) // Лучше всего использовать panic здесь, так как без логгера мы не сможем сообщить об ошибке
	}
}
