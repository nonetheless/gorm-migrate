package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io"
)

var configFile string
var dirName string
var packageName string

func NewRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "gmigrate",
		Short:        "The Migrate generator for Gorm.",
		SilenceUsage: true,
	}
	flags := cmd.PersistentFlags()
	out := cmd.OutOrStdout()
	flags.StringVarP(&configFile, "configFile", "c", "configFile", "gorm migrate config")
	flags.StringVarP(&dirName, "dirName", "d", "dirName", "gorm migrate new version root dir path")
	flags.StringVarP(&packageName, "packageName", "p", "packageName", "your project test_migration package name packageName")
	flags.Parse(args)
	cmdOut := CmdOut{out:out}
	cmd.AddCommand(
		newMigrateCmd(&dirName, &packageName, &cmdOut),
	)
	cmd.AddCommand(
		newStampCmd(&dirName, &cmdOut),
	)

	return cmd
}

type CmdOut struct {
	out io.Writer
}

func (cmdOut *CmdOut) Infof(info string, opts ...interface{}) {
	cmdOut.out.Write([]byte("[INFO] "))
	infoString := fmt.Sprintf(info, opts...)
	cmdOut.out.Write([]byte(infoString))
}

func (cmdOut *CmdOut) Errorf(info string, opts ...interface{}) {
	cmdOut.out.Write([]byte("[ERROR] "))
	infoString := fmt.Sprintf(info, opts...)
	cmdOut.out.Write([]byte(infoString))
}

func (cmdOut *CmdOut) Infoln(info string,opts ...interface{}) {
	cmdOut.out.Write([]byte("[INFO] "))
	cmdOut.out.Write([]byte(info + "\n"))
}

func (cmdOut *CmdOut) Errorln(info string,opts ...interface{}) {
	cmdOut.out.Write([]byte("[ERROR] "))
	cmdOut.out.Write([]byte(info+"\n"))
}
