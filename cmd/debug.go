package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func debugCmd(cfg *rootOpts, args []string) {
	fmt.Printf("%#v\n", cfg)
	viper.Debug()
}

func newDebugCmd(cfg *rootOpts) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "debug",
		Short: "debug",
		Run: func(_ *cobra.Command, args []string) {
			debugCmd(cfg, args)
		},
	}
	return cmd
}
