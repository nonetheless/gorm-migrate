package version

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nonetheless/gorm-migrate/asset"
	"github.com/nonetheless/gorm-migrate/pkg/migrate"
	"testing"
)

func TestAsset(t *testing.T){
	bytes, _ := asset.Asset("template/doc.gotmpl")
	fmt.Println(string(bytes))
}

func testDownGrade(t *testing.T){
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

func testUpgrade(t *testing.T) {
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

func TestUpgradeDownGrade(t *testing.T){
	t.Run("TestUpgradeDownGrade",testUpgrade)
	t.Run("testDownGrade",testDownGrade)

}

type TestCmd struct {

}

func (t *TestCmd) Infof(info string, opts ...interface{}){
	fmt.Printf(info, opts...)
}

func (t *TestCmd) Errorf(info string, opts ...interface{}){
	fmt.Printf(info, opts...)
}

func (t *TestCmd) Infoln(info string, opts ...interface{}){
	fmt.Println(info)
}

func (t *TestCmd) Errorln(info string, opts ...interface{}){
	fmt.Println(info)
}
