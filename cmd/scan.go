package cmd

import (
	"errors"
	usr "os/user"
	"time"

	"github.com/a2-ai/sorting-hat/internal/users"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type scanOpts struct {
	dir string
}

type scanCmd struct {
	cmd  *cobra.Command
	opts *scanOpts
}

func (opts *scanOpts) Set() {
	opts.dir = viper.GetString("dir")
}

func (opts *scanOpts) Validate() error {
	// TODO: check dir exists
	if opts.dir == "" {
		return errors.New("please specify a directory")
	}
	return nil
}

func newScan(cfg *settings, opts *scanOpts, args []string) error {
	// fmt.Printf("%#v\n", cfg)
	startTime := time.Now()
	grp, err := usr.LookupGroup("rstudio-connect")
	if err != nil {
		return err
	}
	potentialUsers, err := users.GetPotentialUsersByDirName(opts.dir, cfg.logger, 10)
	if err != nil {
		return err
	}
	var usersAdded []string
	for _, u := range potentialUsers.Users {
		// we can log the number of groups detected as a sanity check
		// in case there is some issue that is swallowed aroung group lookup,
		// but actually logging what groups the user is part of would get way
		// too noisy
		if lo.Contains(u.GroupIDs, grp.Gid) {
			cfg.logger.Debug("found in rstudio-connect group", "user", u.Username, "n_grps", len(u.GroupIDs))
		} else {
			cfg.logger.Debug("NOT found in rstudio-connect group", "user", u.Username, "n_grps", len(u.GroupIDs))
			err := u.AddToGroupByName("rstudio-connect")
			if err != nil {
				return err
			}
			usersAdded = append(usersAdded, u.Username)
		}
	}
	cfg.logger.Info("ran user check on rstudio-connect group membership",
		"users_added", len(usersAdded),
		"users_checked", len(potentialUsers.Users),
		"duration_secs", time.Since(startTime).Seconds(),
	)
	cfg.logger.Debug("users added", "names", usersAdded)
	return nil
}

func newScanCmd(cfg *settings) *scanCmd {
	scanCmd := &scanCmd{opts: &scanOpts{}}
	cmd := &cobra.Command{
		Use:   "scan",
		Short: "scan",
		PreRunE: func(_ *cobra.Command, args []string) error {
			scanCmd.opts.Set()
			return scanCmd.opts.Validate()
		},
		RunE: func(_ *cobra.Command, args []string) error {
			return newScan(cfg, scanCmd.opts, args)
		},
	}

	cmd.Flags().String("dir", "", "directory for user homes")
	viper.BindPFlag("dir", cmd.Flags().Lookup("dir"))
	scanCmd.cmd = cmd
	return scanCmd
}
