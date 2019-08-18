package version

import (
	"container/list"
	"github.com/jinzhu/gorm"
	api "github.com/nonetheless/gorm-migrate/pkg"
	mig "github.com/nonetheless/gorm-migrate/pkg/migrate"
	param "packageName/param"
	_ "packageName/eu5o4aeza4cz"
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
	if err == nil{
		panic(err)
	}
	return migrate

}