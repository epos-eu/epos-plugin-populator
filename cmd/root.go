package cmd

import (
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var Version = "dev"

func getVersion() string {
	if Version != "" && Version != "dev" {
		return Version
	}

	if info, ok := debug.ReadBuildInfo(); ok && info.Main.Version != "(devel)" {
		return info.Main.Version
	}
	return "dev"
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "epos-plugin-populator",
	Short: "Populate an EPOS Platform environment with plugins for the converter",
	Long:  `Populate an EPOS Platform environment with plugins for the converter`,

	// If no subcommand is provided, show help.
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		}
	},
	Version: getVersion(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
