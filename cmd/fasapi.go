package cmd

import (
	"fmt"
	"github.com/NC-Cj/go-gen/cmd/flags"
	"github.com/spf13/cobra"
)

var fastapiCmd = &cobra.Command{
	Use:     "fastapi",
	Short:   "Scaffold for Quickly Generating Python Language FastAPI Framework",
	Example: "gen fastapi --type=[norm,easy]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("fastapi")
	},
}

func init() {
	rootCmd.AddCommand(fastapiCmd)

	fastapiCmd.Flags().StringVarP(&flags.Name, "name", "n", "", "Project Name")
	fastapiCmd.Flags().BoolVarP(&flags.Db, "db", "", true, "Project with database")

	fastapiCmd.Flags().BoolVarP(&flags.PrismaORM, "prisma", "p", true, "Database ORM and Migration")
	fastapiCmd.Flags().StringVarP(&flags.TypeFlag, "type", "t", "norm", "Template Rule Type")

	fastapiCmd.MarkFlagRequired("name")
	fastapiCmd.MarkFlagRequired("db")

}
