package controller

import ( 
	
	"github.com/gofiber/fiber/v2" 
	"github.com/fiber-gorm/model"
)

func ShowRiset(c *fiber.Ctx) error {
	var riset []model.RisetPenelitian
	model.DB.Find(&riset)
	return c.Status(fiber.StatusOK).JSON(&riset);

}

func CreateRiset(c *fiber.Ctx) error {
	var riset model.RisetPenelitian
	if err := c.BodyParser(&riset); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Create(&riset).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "berhasil",
		"data" : riset,
	})
}

func ShowRisetID(c *fiber.Ctx) error {
	var riset model.RisetPenelitian
	idriset := c.Params("idriset")

	if err := model.DB.Where("idriset",idriset).First(&riset).Error; err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data" : riset,
	})
}

func UpdateRiset(c *fiber.Ctx) error {
	var riset model.RisetPenelitian
	idriset := c.Params("idriset")

	if err := c.BodyParser(&riset);err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := model.DB.Begin().Where("idriset = ?",idriset).Updates(&riset).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "update success",
		"data" : riset,
	})
}


func DeleteRiset(c *fiber.Ctx) error {
	var riset model.RisetPenelitian
	idriset := c.Params("idriset")

	if err := model.DB.Begin().Where("idriset = ?",idriset).Delete(&riset).Error;err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message" : "delete success",
	})
}