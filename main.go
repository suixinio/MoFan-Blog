package main

import (
	"mofan-blog/model"
	"mofan-blog/router"
)

func main() {

	model.InitDB()

	router.InitRouter()
}
