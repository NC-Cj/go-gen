package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gen",
	Long:  `All software has versions. This is gen's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gen Static Site Generator v1.0 -- HEAD")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
