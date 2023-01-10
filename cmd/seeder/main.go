package main

import (
	"fmt"

	"TaskList/config"
	"TaskList/driver"
	"TaskList/pkg/seeds"

	"gorm.io/gorm"
)

func main() {
	config.InitEnv()
	orm := driver.InitGorm()
	// Create Tasks ----------------------------------------------------------------
	taskSeeds := seeds.AllTask()
	run(orm, taskSeeds)

}

func run(orm *gorm.DB, channelSeeds []seeds.Seed) {
	for _, seed := range channelSeeds {
		fmt.Println(seed.Name)
		err := seed.Run(orm)
		if err != nil {
			fmt.Println(seed.Name+" Failed: ", err.Error())
		}
	}
}
