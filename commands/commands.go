package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/sysnet/discovery"
)

var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "discovery",
		Short: "Sample command",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			discovery.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
