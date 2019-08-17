package param

import (
    "fmt"
	api "github.com/nonetheless/gorm-migrate/pkg"
	"sync"
)

var MigrateMap map[string]api.MigrateInterface;

var initFlag sync.Once

func Register(mig api.MigrateInterface) {
	initFlag.Do(func() {
		MigrateMap = make(map[string]api.MigrateInterface)
	})
	//add mig to migrateList
	if _, ok := MigrateMap[mig.PreVersion()]; ok {
		// error for it
		panic(fmt.Errorf("There is some version has same preversion"))
	}
	MigrateMap[mig.PreVersion()] = mig

}
