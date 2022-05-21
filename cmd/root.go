package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bh",
	Short: "GinForBH is an api for bililive handle",
	Run: func(cmd *cobra.Command, args []string) {
		Error(cmd, args, errors.New("unrecognized command"))
	},
}

func Execute() {
	rootCmd.Execute()
}
