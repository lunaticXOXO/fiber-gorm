package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/fiber-gorm/controller"
)


func Start(){

	app := fiber.New()
	api := app.Group("/api")
	
	usertype := api.Group("/usertype")
	users := api.Group("/user")
	riset := api.Group("/riset")
	peneliti := api.Group("/peneliti")

	usertype.Post("/",controller.CreateUserType)
	

	users.Post("/register",controller.CreateUser)
	users.Post("/login",controller.LoginUser)
	users.Post("/logout",controller.Logout)

	
	riset.Get("/",controller.Authenticate(controller.ShowRiset,1))
	riset.Post("/",controller.Authenticate(controller.CreateRiset,1))
	riset.Get("/:idriset",controller.Authenticate(controller.ShowRisetID,1))
	riset.Put("/:idriset",controller.Authenticate(controller.UpdateRiset,1))
	riset.Delete("/:idriset",controller.Authenticate(controller.DeleteRiset,1))
	
	peneliti.Post("/",controller.Authenticate(controller.CreatePeneliti,1))
	peneliti.Get("/",controller.Authenticate(controller.GetPeneliti,1))
	peneliti.Put("/:nidn",controller.Authenticate(controller.UpdatePeneliti,1))
	peneliti.Delete("/:nidn",controller.Authenticate(controller.DeletePeneliti,1))
	
	peneliti.Get("/joins_riset/:nidn",controller.Authenticate(controller.GetJoinRiset,1))
	peneliti.Get("/specified/:nidn",controller.Authenticate(controller.GetSpecifiedColumn,1))
	peneliti.Get("/specified_join/:nidn",controller.Authenticate(controller.GetSpecifiedJoin,1))

	app.Listen(":8080")
}