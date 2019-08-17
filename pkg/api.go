package api

import "github.com/jinzhu/gorm"

type MigrateInterface interface {
	Run(db *gorm.DB) error
	RollBack(db *gorm.DB) error
	Version() string
	PreVersion() string
}

type Option func(MigrateController)

type MigrateController interface {
	Migrate(...Option) error
	Upgrade(...Option) error
	Downgrade(...Option) error
}
