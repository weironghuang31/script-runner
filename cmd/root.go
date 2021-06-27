package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/weironghuang31/script-runner/scripts"
)

const SpecFilename = ".scripts.yml"

var Version = "0.0.0"

var rootCmd = &cobra.Command{
	Use:     "run [flags] [script-name...]",
	Example: "run foo bar",
	Short:   "Run scripts",
	Long:    `Run scripts which is defined in .scripts.yml file`,
	Args:    cobra.MinimumNArgs(1),
	Version: Version,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		dir, _ := cmd.Flags().GetString("dir")
		spec, err := scripts.ParseSpec(dir, SpecFilename)

		if err != nil {
			return err
		}

		return spec.Run(args)
	},
}

func init() {
	rootCmd.Flags().StringP("dir", "d", "", "The working directory. The default value is current working directory.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
