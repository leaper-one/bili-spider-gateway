package cmd

import (
	"fmt"
	"ginForBH/gins"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	var daemon bool
	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start GinForBH",
		Run: func(cmd *cobra.Command, args []string) {
			if daemon {
				command := exec.Command("./bh", "start")
				command.Start()
				fmt.Printf("GinForBH start, [PID] %d running...\n", command.Process.Pid)
				ioutil.WriteFile("GinForBH.lock", []byte(fmt.Sprintf("%d", command.Process.Pid)), 0666)
				daemon = false
				os.Exit(0)
			} else {
				fmt.Println("GinForBH start")
			}
			gins.GinService()
		},
	}
	startCmd.Flags().BoolVarP(&daemon, "deamon", "d", false, "is daemon?")
	rootCmd.AddCommand(startCmd)

}
