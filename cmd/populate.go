package cmd

import (
	"encoding/json"
	"epos-plugin-populator/cmd/internal"
	"epos-plugin-populator/display"
	"fmt"
	"net/url"
	"os"

	"github.com/spf13/cobra"
)

var versionFlag string

var populateCmd = &cobra.Command{
	Use:   "populate [gateway URL] [path to plugins JSON file]",
	Short: "Populate an EPOS Platform environme with converter plugins",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		gatewayURL := args[0]
		pluginsFilePath := args[1]

		baseURL, err := url.Parse(gatewayURL)
		if err != nil {
			panic(fmt.Errorf("error parsing gateway URL: %w", err))
		}

		pluginsFile, err := os.ReadFile(pluginsFilePath)
		if err != nil {
			panic("TODO")
		}

		var plugins []internal.Plugin
		err = json.Unmarshal(pluginsFile, &plugins)
		if err != nil {
			panic("TODO")
		}

		display.Step("Starting population of EPOS Platform environment at '%s', with plugins file '%s'", baseURL.String(), pluginsFilePath)

		err = internal.Populate(*baseURL, plugins, versionFlag)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	populateCmd.Flags().StringVar(&versionFlag, "plugin-version", "", "If set it will override all version of every plugin to the set string")
	rootCmd.AddCommand(populateCmd)
}
