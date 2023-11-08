package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type settings struct {
	// logrus log level
	logger *slog.Logger
}

type rootCmd struct {
	cmd *cobra.Command
	cfg *settings
}

func Execute(version string, args []string) {
	newRootCmd(version).Execute(args)
}

func (cmd *rootCmd) Execute(args []string) {
	cmd.cmd.SetArgs(args)
	if err := cmd.cmd.Execute(); err != nil {
		// if get to this point and don't fatally log in the subcommand,
		// the Usage help will be printed before the error,
		// which may or may not be the desired behavior
		slog.Error(fmt.Sprintf("failed with error: %s", err))
		os.Exit(1)
	}
}

func getSlogLevel(lvl string) (slog.Level, error) {
	switch lvl {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	default:
		return slog.LevelDebug, fmt.Errorf("invalid log level defined: `%s`, must be one of: debug,info,warn,error", lvl)
	}
}

func newRootCmd(version string) *rootCmd {
	root := &rootCmd{cfg: &settings{}}
	cmd := &cobra.Command{
		Use:   "sortinghat",
		Short: "sortinghat is a tool for adding users to local groups as they enter the system",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// need to set the config values here as the viper values
			// will not be processed until Execute, so can't
			// set them in the initializer.
			// If persistentPreRun is used elsewhere, should
			// remember to setGlobalSettings in the initializer
			logLevel, err := getSlogLevel(viper.GetString("loglevel"))
			if err != nil {
				slog.Error(fmt.Sprintf("failed to get log level: %s", err))
				os.Exit(1)
			}
			logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: logLevel,
			}))
			root.cfg.logger = logger
		},
	}
	cmd.Version = version
	// without this, the default version is like `cmd version <version>` so this
	// will just print the version for simpler parsing
	cmd.SetVersionTemplate(`{{printf "%s\n" .Version}}`)
	cmd.PersistentFlags().String("loglevel", "info", "log level")
	viper.BindPFlag("loglevel", cmd.PersistentFlags().Lookup("loglevel"))
	cmd.AddCommand(newDebugCmd(root.cfg))
	cmd.AddCommand(newWatchCmd(root.cfg).cmd)
	root.cmd = cmd
	return root
}
