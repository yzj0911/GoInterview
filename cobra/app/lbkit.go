package app

import (
	"execlt1/cobra/app/cmd"
	"github.com/lithammer/dedent"
	"github.com/spf13/cobra"
	"io"
	"os"
)

func Run() error {
	kitCmd := NewkitCommand(os.Stdin, os.Stdout, os.Stderr)
	return kitCmd.Execute()
}
func NewkitCommand(in io.Reader, out, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:           "lbkit",
		Short:         "lbkit:linkBackup 的工具箱",
		Long:          dedent.Dedent(``),
		SilenceErrors: true,
		SilenceUsage:  true,
	}
	cmds.ResetFlags()

	cmds.AddCommand(cmd.NewCmdMysql(out, err))
	return cmds
}
