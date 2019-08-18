# gorm-migrate
A codegen tool to generate gorm versiond migration code, you can create a gorm migration with gorm-migrate client, and you can use the migration code by import it.

## Migrate codegen tool
### Install with curl
```bash

```
### Install and build
```bash
go get -u github.com/nonetheless/gorm-migrate
```

### Create a migration and create a new version
You can easily use this tool to create you gorm migration
```bash
gorm-migrate 
The Migrate generator for Gorm.

Usage:
  gmigrate [command]

Available Commands:
  help        Help about any command
  migrate     create a new version
  stamp       check test_migration version relation

Flags:
  -c, --configFile string    gorm migrate config (default "configFile")
  -d, --dirName string       gorm migrate new version root dir path (default "dirName")
  -h, --help                 help for gmigrate
  -p, --packageName string   your project test_migration package name packageName (default "packageName")

# you can create migration with this cmd
gorm-migrate migrate -d $DIR -p $PACKAGENAME
......

# you can check the migration with this cmd
gorm-migrate stamp -d $DIR
[INFO] Stamp test_migration version:
[INFO]                --------->     xlq9lndwjv9j
[INFO] xlq9lndwjv9j   --------->     phma9g1znhpm
[INFO] phma9g1znhpm   --------->     p1buif9wgmig
[INFO] p1buif9wgmig   --------->     v3c06skwrur9

```
## Use Migration 
### Write your own version migration function
With the tool you can get the code template like:
```go
func run(db *gorm.DB) error {
	//TODO add version update sql
	return nil
}

func rollBack(db *gorm.DB) error {
	//TODO add version rollback function
	return nil
}
```
You must write your own migration function with gorm.DB

### How to use your migration
You can use your migration code with import code, this is a upgrade database example:
```go
pacakge main
// main.go
import (
	"github.com/google/martian/log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
    //this is your own migration pacakge
	migration "github.com/nonetheless/gorm-migrate/pkg/migrate"
)

func main(){
    	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True")
    	if err != nil {
    		t.Fatal(err)
    	}
    	defer db.Close()
    	update, err:= migration.CreateMigration(db)
    	err = update.Upgrade(migrate.WithCmdOut(log))
    	if err != nil {
    		t.Fatal(err)
    	}
}



```

