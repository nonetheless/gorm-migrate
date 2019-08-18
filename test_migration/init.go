package version

import (
	"container/list"
	"github.com/jinzhu/gorm"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
	_ "github.com/nonetheless/gorm-migrate/test_migration/p1buif9wgmig"
	"github.com/nonetheless/gorm-migrate/test_migration/param"
	_ "github.com/nonetheless/gorm-migrate/test_migration/phma9g1znhpm"
	_ "github.com/nonetheless/gorm-migrate/test_migration/v3c06skwrur9"
	_ "github.com/nonetheless/gorm-migrate/test_migration/xlq9lndwjv9j"
)

func CreateMigration(db *gorm.DB) (api.MigrateController, error) {
	migList := list.New()
	migList.PushFront(nil)
	migList.PushBack(nil)
	head := migList.Front()
	version := ""
	for {
		if mig, ok := param.MigrateMap[version]; ok {
			version = mig.Version()
			migList.InsertAfter(mig, head)
			head = head.Next()
		} else {
			break
		}
	}
	migrate, err := mig.NewMigration(db, migList)
	if err != nil {
		return nil, err
	}
	return migrate, nil

}
