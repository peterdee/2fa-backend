package utilities

import "github.com/gofiber/fiber/v2"

type ResponsePayloadStruct struct {
	Context *fiber.Ctx
	Data    interface{}
	Info    string
	Status  int
}
