package main

import (
	"cms/internal/config"
	"cms/internal/migration"
)

func main() {
	err := config.InitConfig()
	if err != nil {
		println(err.Error())
		return
	}

	err = migration.Start()
	if err != nil {
		println("Migration: ", err.Error())
	} else {
		println("Migration: Done...")
	}
}
