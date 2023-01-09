package main

import (
	"fmt"

	"TaskList/config"
	"TaskList/pkg/seeds"

	"gorm.io/gorm"
)

func main() {
	config.InitEnv()
	//orm := driver.InitGorm()

}

func run(orm *gorm.DB, channelSeeds []seeds.Seed) {
	for _, seed := range channelSeeds {
		fmt.Println(seed.Name)
		err := seed.Run(orm)
		if err != nil {
			fmt.Println(seed.Name + " Failed")
			fmt.Println(err.Error())
		}
	}
}
