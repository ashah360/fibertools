package fibertools

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

// Error handler that is passed into `fiber.New()`. Set x-debug in request header to see stack trace of error.
func ErrorHandler(c *fiber.Ctx, err error) error {
	isDebug := GetHeader(c, "x-debug")

	richErr, ok := err.(*RichError)
	if !ok {
		richErr = NewError(err)
	}

	if isDebug == "true" {
		return c.Status(richErr.Code).JSON(fiber.Map{
			"msg":   richErr.Message,
			"code":  richErr.Code,
			"error": richErr.StackTrace(),
		})
	}

	log.Println(richErr.Message)

	return Message(c, richErr.Code, richErr.Message)
}
