package cmd

import (
	"fmt"
	"os"

	"oeis/cli/oeis"
	"oeis/pkg/lib/input"
	"oeis/pkg/lib/output"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "oeis [ARGS]",
	Args:  cobra.MinimumNArgs(1),
	Short: "cli to query oeis",
	Long:  "CLI built on golang to query The On-Line Encyclopedia of Integer Sequences",
	Run: func(cmd *cobra.Command, args []string) {
		inputData := input.Format(args)
		err := oeis.SearchAPI(inputData)
		if err != nil {
			output.PrintError(err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
