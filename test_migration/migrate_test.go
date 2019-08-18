package version

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nonetheless/gorm-migrate/pkg/migrate"
	"testing"
)

func TestDownGrade(t *testing.T){
	db, err := gorm.Open("mysql", "root:66166161@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	update, err:= CreateMigration(db)
	if err != nil {
		t.Fatal(err)
	}
	testCmdOut := TestCmd{}
	err = update.Downgrade(migrate.WithCmdOut(&testCmdOut))
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpgrade(t *testing.T) {
	db, err := gorm.Open("mysql", "root:66166161@tcp(127.0.0.1:3306)/demo?charset=utf8&parseTime=True")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	update, err:= CreateMigration(db)
	if err != nil {
		t.Fatal(err)
	}
	testCmdOut := TestCmd{}
	err = update.Upgrade(migrate.WithCmdOut(&testCmdOut))
	if err != nil {
		t.Fatal(err)
	}

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
