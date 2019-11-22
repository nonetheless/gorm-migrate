package api

import "github.com/jinzhu/gorm"

type MigrateInterface interface {
	Run(db *gorm.DB) error
	RollBack(db *gorm.DB) error
	Version() string
	PreVersion() string
	Printer
}

type Option func(MigrateController)

type MigrateController interface {
	Migrate(...Option) error
	Upgrade(...Option) error
	Downgrade(...Option) error
	Stamp(...Option) error
	GetDbVersion(...Option) string
}

type MigrateOut interface {
	Infoln(...interface{})
	Errorln(...interface{})
	Infof(string, ...interface{})
	Errorf(string, ...interface{})

}

type Printer interface {
	Printf(MigrateOut)
	RPrintf(out MigrateOut)
}

