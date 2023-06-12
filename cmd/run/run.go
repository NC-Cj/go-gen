package run

import (
	"fmt"
	"github.com/NC-Cj/go-gen/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strconv"
)

const normRepositoryAddress = "https://github.com/NC-Cj/gen-fastapi-norm.git"
const normRepositoryName = "/gen-fastapi-norm"

const easyRepositoryAddress = "https://github.com/NC-Cj/gen-fastapi-easy.git"
const easyRepositoryName = "/gen-fastapi-easy"

func FastapiRun() func(cmd *cobra.Command, args []string) {
	var obj pkg.CloneObject
	var type_ string

	return func(cmd *cobra.Command, args []string) {
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Name == "name" {
				obj.ProjectName = f.Value.String()
			} else if f.Name == "prisma" {
				obj.Prisma, _ = strconv.ParseBool(f.Value.String())
			} else if f.Name == "type" {
				type_ = f.Value.String()
			}
		})

		if type_ == "norm" {
			obj.RepositoryAddress = normRepositoryAddress
			obj.RepositoryName = normRepositoryName
		} else if type_ == "easy" {
			obj.RepositoryAddress = easyRepositoryAddress
			obj.RepositoryName = easyRepositoryName
		}
		fmt.Println(obj)
		obj.Generate()
	}
}
