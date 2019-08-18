package migrate

import (
	"github.com/jinzhu/gorm"
	param "packageName/param"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
)

const (
	version = "eu5o4aeza4cz"
	preversion = ""
)

func run(db *gorm.DB) error {
	//TODO add version update sql
	return nil
}

func rollBack(db *gorm.DB) error {
	//TODO add version rollback function
	return nil
}

func NewMigrate() api.MigrateInterface{
	migrateVersion := mig.NewMigrateVersion(version,preversion,run,rollBack)
	return migrateVersion
}

func init(){
    param.Register(NewMigrate())
}
