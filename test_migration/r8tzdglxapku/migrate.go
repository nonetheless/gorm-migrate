package migrate

import (
	"github.com/jinzhu/gorm"
	param "github.com/nonetheless/gorm-migrate/test_migration/param"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
)

const (
	version = "r8tzdglxapku"
	preversion = "v3c06skwrur9"
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
