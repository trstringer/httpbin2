/*
Copyright © 2023 Thomas Stringer <thomas@trstringer.com>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/trstringer/httpbin2/version"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Application version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
