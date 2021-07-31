package mysql

import (
	"github.com/spf13/cobra"
	"io"
)

func NewMysqlBackup(out, err io.Writer) *cobra.Command {
	var (
		op string
	)
	mysqlBackup := &cobra.Command{
		Use: "mysql",
		RunE: func(cmd *cobra.Command, args []string) error {
			switch op {
			case "aa":
				return nil
			}
			return nil
		},
	}
	mysqlBackup.Flags().StringVarP(&op, "op", "", "", "Mysql ")
	return mysqlBackup
}
