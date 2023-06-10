package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var fastapiCmd = &cobra.Command{
	Use:     "fastapi",
	Short:   "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fastapi")
	},
}

func init() {
	rootCmd.AddCommand(fastapiCmd)
}
