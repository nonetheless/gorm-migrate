package migrate

import (
	"github.com/jinzhu/gorm"
	api "github.com/nonetheless/gorm-migrate/pkg"
)

type GormVersion struct {
	gorm.Model
	Version string `gorm:"type:varchar(255)"`
}

func (GormVersion) TableName() string {
	return "gorm_version"
}

type MigrateVersion struct {
	VersionValue    string
	PreVersionValue string
	RunFunc         func(db *gorm.DB) error
	RollbackFunc    func(db *gorm.DB) error
}

func NewMigrateVersion(version, preVersion string, runFunc, rolbackFunc func(db *gorm.DB) error) api.MigrateInterface {
	migrateVersion := MigrateVersion{
		VersionValue:    version,
		PreVersionValue: preVersion,
		RunFunc:         runFunc,
		RollbackFunc:    rolbackFunc,
	}
	return &migrateVersion
}

func (v *MigrateVersion) Run(db *gorm.DB) error {
	return v.RunFunc(db)
}

func (v *MigrateVersion) RollBack(db *gorm.DB) error {
	return v.RollbackFunc(db)
}

func (v *MigrateVersion) Version() string {
	return v.VersionValue
}

func (v *MigrateVersion) PreVersion() string {
	return v.PreVersionValue
}

func (v *MigrateVersion) Printf(out api.MigrateOut){
	tempVersion := v.PreVersion()
	if tempVersion == "" {
		tempVersion = "            "
	}
	out.Infof(tempVersion + "   --------->     " + v.Version() + "\n")
}

func (v *MigrateVersion) RPrintf(out api.MigrateOut){
	tempVersion := v.PreVersion()
	if tempVersion == "" {
		tempVersion = "            "
	}
	out.Infof(tempVersion + "   <---------     " + v.Version() + "\n")
}