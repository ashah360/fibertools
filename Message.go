package fibertools

import "github.com/gofiber/fiber/v2"

func Message(c *fiber.Ctx, code int, msg string) error {
	return c.Status(code).JSON(fiber.Map{
		"code": code,
		"msg": msg,
	})
}
