package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/default_username/default_app_name/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "default_app_name"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
