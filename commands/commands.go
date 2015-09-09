package commands

import (
	"github.com/spf13/cobra"

	"github.com/default_username/default_app_name/sample"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "sample",
		Short: "Sample command",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			sample.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
