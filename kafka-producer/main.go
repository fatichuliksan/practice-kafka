package main

import (
	"kafka-example/api"
	"kafka-example/helper"
	"kafka-example/helper/kafka"
	"kafka-example/helper/response"
	"kafka-example/helper/validator"
	"kafka-example/helper/viper"
	"log"

	"github.com/labstack/echo"
)

// App ...
type App struct {
	config viper.Config
}

var app App

func init() {
	config := viper.NewViper()
	app = App{config: config}
	if config.GetBool(`app.debug`) {
		log.Println("Service RUN on DEBUG mode - HOST: " + config.GetString("app.host"))
	}
}

func main() {
	e := echo.New()

	// kafka inisiate
	kafkaHelper := kafka.KafkaHelper{
		Host: app.config.GetString(`kafka.host`),
	}
	helper := helper.NewHelper{
		Response:  response.ResponseHelper{},
		Config:    app.config,
		Validator: validator.NewValidator(),
		Kafka:     kafkaHelper,
	}
	router := api.NewAPI{
		E:      e,
		Helper: helper,
	}
	router.Register()
	e.Start(app.config.GetString(`app.host`))
}
