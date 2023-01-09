//go:build wireinject
// +build wireinject

package api

import (
	"TaskList/driver"

	"github.com/google/wire"
)

var gormSet = wire.NewSet(driver.InitGorm)
