package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
	param "github.com/nonetheless/gorm-migrate/test_migration/param"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
)

const (
	version = "xlq9lndwjv9j"
	preversion = ""
)

func run(db *gorm.DB) error {
	//TODO add version update sql
	fmt.Println("fist create")
	return nil
}

func rollBack(db *gorm.DB) error {
	//TODO add version rollback function
	fmt.Println("fist rollback")
	return nil
}

func NewMigrate() api.MigrateInterface{
	migrateVersion := mig.NewMigrateVersion(version,preversion,run,rollBack)
	return migrateVersion
}

func init(){
    param.Register(NewMigrate())
}
