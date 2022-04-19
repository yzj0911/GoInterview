package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(fuck)
}

var fuck = &cobra.Command{
	Use:   "fuck",
	Short: "fuck is not good say",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		fmt.Println(cmd)
		fmt.Println("fuc")
	},
}
