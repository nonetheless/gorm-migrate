package main

import "github.com/nonetheless/gorm-migrate/pkg/migrate"

func main() {
	mig, err := migrate.NewMigrationToInit()
	if err != nil {
		panic(err)
	}
	mig.Migrate(migrate.WithDirPath("./migration"), migrate.WithPackageName("github.com/nonetheless/gorm-migrate/migration"))
}
