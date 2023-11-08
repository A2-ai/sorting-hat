package cmd

import (
	"fmt"
	usr "os/user"
	"time"

	"github.com/a2-ai/sorting-hat/internal/users"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

func debugCmd(cfg *settings, args []string) {
	// fmt.Printf("%#v\n", cfg)
	grp, err := usr.LookupGroup("rstudio-connect")
	if err != nil {
		cfg.logger.Error(err.Error())
	}
	potentialUsers, err := users.GetPotentialUsersByDirName("/home", cfg.logger, 10)
	if err != nil {
		cfg.logger.Error(err.Error())
	}
	for _, u := range potentialUsers.Users {
		if lo.Contains(u.GroupIDs, grp.Gid) {
			cfg.logger.Info("found in rstudio-connect group", "user", u.Username, "groups", u.GroupIDs)
		} else {
			cfg.logger.Info("NOT found in rstudio-connect group", "user", u.Username, "groups", u.GroupIDs)
			start_time := time.Now()
			err := u.AddToGroupByName("rstudio-connect")
			fmt.Println("time elapsed:", time.Since(start_time))
			if err != nil {
				cfg.logger.Error("error adding user by name", "error", err.Error())
			}
		}
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
