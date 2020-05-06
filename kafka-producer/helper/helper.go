package helper

import (
	"kafka-example/helper/jwt"
	"kafka-example/helper/kafka"
	"kafka-example/helper/response"
	"kafka-example/helper/validator"
	"kafka-example/helper/viper"
)

// NewHelper ...
type NewHelper struct {
	Response  response.ResponseHelper
	Validator *validator.Validator
	Config    viper.Config
	Jwt       jwt.JwtHelper
	Kafka     kafka.KafkaHelper
}
