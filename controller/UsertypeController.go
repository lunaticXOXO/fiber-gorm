package controller

import (

	"github.com/gofiber/fiber/v2"
	"github.com/fiber-gorm/model"
)

func CreateUserType(c *fiber.Ctx) error	{
	var usertype model.Usertype

	if err := c.BodyParser(&usertype);err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Create(&usertype).Error;err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "add usertype success",
		"data" : usertype,
	})

}