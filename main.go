package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/sysnet/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "sysnet"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
