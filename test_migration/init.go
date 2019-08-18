package version

import (
	"container/list"
	"github.com/jinzhu/gorm"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
	param "github.com/nonetheless/gorm-migrate/test_migration/param"
	_ "github.com/nonetheless/gorm-migrate/test_migration/xlq9lndwjv9j"
	_ "github.com/nonetheless/gorm-migrate/test_migration/phma9g1znhpm"
	_ "github.com/nonetheless/gorm-migrate/test_migration/p1buif9wgmig"
)

func CreateMigration(db *gorm.DB) api.MigrateController{
	migList := list.New()
	migList.PushFront(nil)
	migList.PushBack(nil)
	head := migList.Front()
	version := ""
	for{
		if mig,ok := param.MigrateMap[version];ok{
			version = mig.Version()
			migList.InsertAfter(mig, head)
			head = head.Next()
		} else {
			break
		}
	}
	migrate, err:= mig.NewMigration(db, migList)
	if err != nil{
		panic(err)
	}
	return migrate

}