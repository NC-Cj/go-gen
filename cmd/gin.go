package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var ginCmd = &cobra.Command{
	Use:     "gin",
	Short:   "",
	Example: "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ginCmd")
	},
}

func init() {
	rootCmd.AddCommand(ginCmd)
}
