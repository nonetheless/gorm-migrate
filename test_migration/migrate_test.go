package version

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nonetheless/gorm-migrate/pkg/migrate"
	"testing"
)


func TestUpgrade(t *testing.T) {
	db, err := gorm.Open("mysql", "root:66166161@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	update := CreateMigration(db)
	testCmdOut := TestCmd{}
	update.Upgrade(migrate.WithCmdOut(&testCmdOut))
}

type TestCmd struct {

}

func (t *TestCmd) Infof(info string, opts ...interface{}){
	fmt.Printf(info, opts...)
}

func (t *TestCmd) Errorf(info string, opts ...interface{}){
	fmt.Printf(info, opts...)
}

func (t *TestCmd) Infoln(info string){
	fmt.Println(info)
}

func (t *TestCmd) Errorln(info string){
	fmt.Println(info)
}
