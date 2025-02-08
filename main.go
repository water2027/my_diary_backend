package main

import (
	"log"
	"my_diary/config"
	"my_diary/database"
	"my_diary/router"
	"my_diary/utils"
)


func main() {
	err := utils.InitLog()
	if err != nil {
		log.Panic(err)
		return
	}
	
	err = config.InitConfig()
	if err != nil {
		log.Panic(err)
		return
	}
	
	err = database.InitDatabase()
	if err != nil {
		log.Panic(err)
		return
	}
	
	r := router.RouterHelper()
	err = r.Run(":8080")
	if err != nil {
		log.Panic(err)
		return
	}
}