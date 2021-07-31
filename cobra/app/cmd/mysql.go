package cmd

import (
	"execlt1/cobra/app/cmd/mysql"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func NewCmdMysql(out, err io.Writer) *cobra.Command {
	rootPath := ".../cmd/scripts"
	Mysqlcmd := &cobra.Command{
		Use:           "mysql",
		Short:         "mysql ",
		Long:          dedent.Dedent(``),
		SilenceUsage:  true,
		SilenceErrors: true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return os.Setenv("LB_ROOT", rootPath)
		},
	}
	Mysqlcmd.ResetFlags()

	Mysqlcmd.AddCommand(mysql.NewMysqlBackup(out, err))
	return Mysqlcmd
}
