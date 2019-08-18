package cmd

import (
	api "github.com/nonetheless/gorm-migrate/pkg"
	"github.com/nonetheless/gorm-migrate/pkg/migrate"
	"github.com/spf13/cobra"
)

const servDesc = `
This command manage GORM test_migration.

you can create a new version by :

    $ gmigrate migrate -d dir -p pacakgeName

`

type MigrateCmd struct {
	dirPath     string
	packageName string
	out         api.MigrateOut
}

func newMigrateCmd(dirPath *string, packageName *string, out api.MigrateOut) *cobra.Command {
	cc := &MigrateCmd{
		dirPath:     *dirPath,
		packageName: *packageName,
		out:         out,
	}

	cmd := &cobra.Command{
		Use:   "migrate",
		Short: "create a new version",
		Long:  servDesc,
		RunE: func(cmd *cobra.Command, args []string) error {

			return cc.run()
		},
	}

	return cmd
}

func (cmd *MigrateCmd) run() error {
	mig, err := migrate.NewMigrationToInit()
	if err != nil {
		cmd.out.Errorln(err.Error())
		return err
	}
	if cmd.dirPath == ""{
		cmd.out.Errorln("Must input dir path by flag <-d>")
		return nil
	}
	if cmd.packageName == ""{
		cmd.out.Errorln("Must input package name by flag <-p>")
		return nil
	}
	err = mig.Migrate(migrate.WithDirPath(cmd.dirPath), migrate.WithPackageName(cmd.packageName), migrate.WithCmdOut(cmd.out))
	if err != nil {
		cmd.out.Errorln(err.Error())
		return err
	}
	return err
}
