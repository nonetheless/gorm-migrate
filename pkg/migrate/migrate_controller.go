package migrate

import (
	"container/list"
	"fmt"
	"github.com/jinzhu/gorm"
	api "github.com/nonetheless/gorm-migrate/pkg"
	"github.com/nonetheless/gorm-migrate/pkg/rander"
	"reflect"
	"strings"
)

type Migrate struct {
	migrateList  *list.List
	dbVersion    *list.Element
	db           *gorm.DB
	now          *list.Element
	rollbackFlag bool
	packageName  string
	dirPath      string
}

func NewMigration(db *gorm.DB, migrateList *list.List) (api.MigrateController, error) {
	//TODO: create var
	version := GormVersion{}
	err := db.First(version).Error
	var dbVersion *list.Element;
	if err == nil {
		point := migrateList.Front()
		for {
			point = point.Next()
			if point == migrateList.Back() {
				break
			}
			if task, ok := point.Value.(api.MigrateInterface); ok {
				if task.Version() == version.Version {
					dbVersion = point
					break
				}
			}
		}
	}
	migrate := Migrate{
		migrateList: migrateList,
		db:          db,
		dbVersion:   dbVersion,
	}
	return &migrate, nil
}

func NewMigrationToInit() (api.MigrateController, error){
	migrate := Migrate{
	}
	return &migrate, nil
}

func (mig *Migrate) Migrate(opts ...api.Option) error {
	for _, opt := range opts{
		opt(mig)
	}

	err := rander.RanderTemplate(mig.dirPath, mig.packageName)
	if err != nil {
		return err
	}
	return nil
}

func (mig *Migrate) Upgrade(opts ...api.Option) (err error) {
	for _, opt := range opts {
		opt(mig)
	}
	if mig.dbVersion == nil {
		head := mig.migrateList.Front() //null
		mig.dbVersion = head.Next()
	}
	if mig.now == nil {
		mig.now = mig.dbVersion
	}
	back := mig.migrateList.Back() //null
	for {
		if mig.now == back {
			break
		}
		if task, ok := mig.now.Value.(api.MigrateInterface); ok {
			err = task.Run(mig.db)
			if err != nil {
				//rollback
				mig.rollback()
				return err
			}
			mig.now = mig.now.Next()
			mig.updateVersion(task.Version())
		} else {
			objType := reflect.TypeOf(mig.now.Value)
			return fmt.Errorf("Migration task can't change to MigrateInterface: %v", objType.Name())
		}
	}
	return nil
}

func (mig *Migrate) Downgrade(opts ...api.Option) (err error) {
	// just downgrage one version
	for _, opt := range opts {
		opt(mig)
	}
	apiCall := true
	if mig.rollbackFlag == true {
		mig.rollbackFlag = false
		apiCall = false
	}
	if mig.dbVersion == nil {
		head := mig.migrateList.Front() //null
		mig.dbVersion = head.Next()
	}

	if mig.now == nil {
		mig.now = mig.dbVersion
	}
	head := mig.migrateList.Front()
	if mig.now.Prev() == head {
		// can't downgrade
		return fmt.Errorf("Now it's head can't downgrade")
	} else {
		// run rollback
		if task, ok := mig.now.Prev().Value.(api.MigrateInterface); ok {
			err = task.RollBack(mig.db)
			if err != nil {
				//rollback
				if apiCall {
					task.Run(mig.db)
				} else {
					return err
				}
			}
			mig.now = mig.now.Prev()
			mig.updateVersion(task.Version())
		} else {
			objType := reflect.TypeOf(mig.now.Value)
			return fmt.Errorf("Migration task can't change to MigrateInterface: %v", objType.Name())
		}
	}
	return nil
}

func rollbackFlag(migInterface api.MigrateController) {
	migrate, ok := migInterface.(*Migrate)
	if ok {
		migrate.rollbackFlag = true
	}
}

func WithPackageName(packageName string) api.Option {
	return func(migInterface api.MigrateController) {
		migrate, ok := migInterface.(*Migrate)
		if ok {
			migrate.packageName = packageName
		}
	}
}

func WithDirPath(dirName string) api.Option {
	return func(migInterface api.MigrateController) {
		migrate, ok := migInterface.(*Migrate)
		if ok {
			migrate.dirPath = dirName
		}
	}
}

func (mig *Migrate) updateVersion(version string) error {
	versionNow := GormVersion{}
	err := mig.db.First(versionNow).Error
	if err != nil {
		//update
		versionNow.Version = version
		err = mig.db.Save(versionNow).Error
		if err != nil {
			return err
		}
	} else {
		//create
		versionNow.Version = version
		err = mig.db.Create(versionNow).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (mig *Migrate) rollback() (err error) {
	if mig.dbVersion == nil {
		for {
			err := mig.Downgrade(rollbackFlag)
			if err != nil {
				if strings.Contains(err.Error(), "head can't downgrade") {
					// end of
					break
				}
				return err
			}
		}

	} else {
		for {
			if mig.now == mig.dbVersion {
				break
			}
			mig.Downgrade()
		}
	}
	return nil
}
