package cmd

import (
	"github.com/NC-Cj/go-gen/flags"
	"github.com/NC-Cj/go-gen/run"
	"github.com/spf13/cobra"
)

var fastapiCmd = &cobra.Command{
	Use:     "fastapi",
	Short:   "Scaffold for Quickly Generating Python Language FastAPI Framework",
	Example: "gen fastapi --type=[norm,easy]",
	Run:     run.Fastapi,
}

func init() {
	fastapiCmd.Flags().StringVarP(&flags.Name, "name", "n", "", "Project Name")
	fastapiCmd.Flags().BoolVarP(&flags.PrismaORM, "prisma", "p", true, "Database ORM and Migration")
	fastapiCmd.Flags().StringVarP(&flags.TypeFlag, "type", "t", "norm", "Template Rule Type")

	fastapiCmd.MarkFlagRequired("name")

	rootCmd.AddCommand(fastapiCmd)

}
