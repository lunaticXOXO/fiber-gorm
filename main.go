package main

import (
	"github.com/fiber-gorm/model"
	"github.com/fiber-gorm/route"

)

func main(){

	model.Connect()
	route.Start()
}