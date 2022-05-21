package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	rootCmd.AddCommand(stopCmd)
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop GinForBH",
	Run: func(cmd *cobra.Command, args []string) {
		strb, _ := ioutil.ReadFile("GinForBH.lock")
		command := exec.Command("tskill", string(strb))
		command.Start()
		os.Remove("GinForBH.lock")
		println("bh stop")
	},
}
