package main

import (
	"log"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/onimarufl/neversitup-test-template/handler"
	"github.com/onimarufl/neversitup-test-template/repository"
	"github.com/onimarufl/neversitup-test-template/service"
	"github.com/spf13/viper"
)

func init() {

	initViper()

}

func main() {
	service := service.NewService(repository.NewRepository())
	handler := handler.NewHandler(service)

	app := fiber.New()

	app.Get("/alive", func(c fiber.Ctx) error {
		return c.JSON("alive")
	})

	app.Get("/getUserById", handler.HandlerGetUser)

	app.Listen(":" + viper.GetString("app.port"))

}

func initViper() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("cannot read in viper config:%s", err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
