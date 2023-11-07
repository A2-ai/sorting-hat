package cmd

import (
	"github.com/a2-ai/sorting-hat/internal/users"
	"github.com/spf13/cobra"
)

func debugCmd(cfg *settings, args []string) {
	// fmt.Printf("%#v\n", cfg)
	potentialUsers, err := users.GetPotentialUsersByDirName("/home", cfg.logger, 10)
	if err != nil {
		cfg.logger.Error(err.Error())
	}
	for _, u := range potentialUsers.Users {
		cfg.logger.Info("userinfo", "user", u.Username, "groups", u.GroupIDs)
	}
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
