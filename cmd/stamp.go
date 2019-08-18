package cmd

import (
	api "github.com/nonetheless/gorm-migrate/pkg"
	"github.com/nonetheless/gorm-migrate/pkg/migrate"
	"github.com/spf13/cobra"
)

const stampDesc = `
This command manage GORM test_migration.

you can check test_migration version relation by :

    $ gmigrate stamp -d dir

`

type StampCmd struct {
	dirPath     string
	out         api.MigrateOut
}

func newStampCmd(dirPath *string, out api.MigrateOut) *cobra.Command {
	cc := &StampCmd{
		dirPath:     *dirPath,
		out:         out,
	}

	cmd := &cobra.Command{
		Use:   "stamp",
		Short: "check test_migration version relation",
		Long:  stampDesc,
		RunE: func(cmd *cobra.Command, args []string) error {

			return cc.run()
		},
	}

	return cmd
}

func (cmd *StampCmd) run() error {
	if cmd.dirPath == ""{
		cmd.out.Errorln("Must input dir path by flag <-d>")
		return nil
	}
	mig, err := migrate.NewMigrationToInit()
	if err != nil {
		cmd.out.Errorln(err.Error())
		return err
	}
	err = mig.Stamp(migrate.WithDirPath(cmd.dirPath),  migrate.WithCmdOut(cmd.out))
	if err != nil {
		cmd.out.Errorln(err.Error())
		return err
	}
	return err
}
