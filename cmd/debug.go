package cmd

import (
	"github.com/spf13/cobra"
)

func debugCmd(cfg *settings, args []string) {

}

func newDebugCmd(cfg *settings) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "debug",
		Run: func(_ *cobra.Command, args []string) {
			debugCmd(cfg, args)
		},
	}
	return cmd
}
