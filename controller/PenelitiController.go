package controller

import (
	
	"github.com/fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
)

func CreatePeneliti(c *fiber.Ctx) error {
	var peneliti model.Peneliti
	if err := c.BodyParser(&peneliti);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Create(&peneliti).Error;err!= nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "success",
		"data" : peneliti,
	})

}

func GetPeneliti(c *fiber.Ctx) error{

	var peneliti []model.Peneliti
	if err := model.DB.Find(&peneliti).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(&peneliti)
}

 


func UpdatePeneliti(c *fiber.Ctx) error {
	var peneliti model.Peneliti
	nidn := c.Params("nidn")

	if err := c.BodyParser(&peneliti); err != nil{
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	if err := model.DB.Begin().Where("nidn",nidn).Updates(&peneliti).Error;err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "update success",
		"data" 	  : peneliti,
	})

}

func DeletePeneliti(c *fiber.Ctx) error {
	var peneliti model.Peneliti
	nidn := c.Params("nidn")

	if err := model.DB.Begin().Where("nidn",nidn).Delete(&peneliti).Error;err != nil{
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "delete success",
	})

}