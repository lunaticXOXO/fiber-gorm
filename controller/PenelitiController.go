package controller

import (
	
	"github.com/fiber-gorm/model"
	"github.com/gofiber/fiber/v2"
	//"fmt"
	//"gorm.io/gorm"
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

func GetJoinRiset(c *fiber.Ctx) error {
	nidn := c.Params("nidn")
	var peneliti model.Peneliti
	if err := model.DB.Preload("RisetPenelitian").Where("nidn",nidn).Find(&peneliti).Error; err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data" : peneliti,
	})

}

func GetSpecifiedColumn(c *fiber.Ctx) error {
	nidn := c.Params("nidn")
	var peneliti model.Peneliti

	if err := model.DB.Select("nidn,nama,telephone").Where("nidn",nidn).Find(&peneliti).Error; err != nil{
		return c.Status(400).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data" : peneliti,
	})
}

func GetSpecifiedJoin(c *fiber.Ctx) error {
	var peneliti model.Peneliti
	nidn := c.Params("nidn")
	
	model.DB = model.DB.Debug()

	query := model.DB.Preload("RisetPenelitian").Where("nidn",nidn).Find(&peneliti)

	if query.Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": query.Error.Error(),
		})
	}

	riset_obj :=  &model.RisetPenelitian{
		Idriset: peneliti.RisetPenelitian.Idriset,
		Judul: peneliti.RisetPenelitian.Judul,
	}

	peneliti_obj := &model.Peneliti{
		Nidn: peneliti.Nidn,
		Nama: peneliti.Nama,
		RisetPenelitian: *riset_obj,

	}

	return c.Status(200).JSON(fiber.Map{
		"data": peneliti_obj,
	})
}
