package controller

import (

	"github.com/fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)


func CreateUser(c *fiber.Ctx) error {
	var users model.Users	
	if err := c.BodyParser(&users); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	if err := model.DB.Create(&users).Error; err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "register success",
		"data" : users,
	})

}	

func LoginUser(c *fiber.Ctx) error {
	var users model.Users
	if err := c.BodyParser(&users);err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error" : err.Error(),
		})
	}
	if err := model.DB.Begin().Where("username = ? AND password = ?",users.Username,users.Password).First(&users).Error;err != nil{
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
		
	err,token := GenerateToken(c,users)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}	

	return c.JSON(fiber.Map{"token": token,"message" : "login success"})
}


func Logout(c *fiber.Ctx) error{
	ResetToken(c)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "logout success",
	})

}