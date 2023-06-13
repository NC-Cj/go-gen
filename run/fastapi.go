package run

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"strconv"
)

func Fastapi(cmd *cobra.Command, args []string) {
	var obj CloneObject
	var type_ string

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
		obj.RepositoryAddress = confVariable["fastapi"]["normAddress"]
		obj.RepositoryName = confVariable["fastapi"]["normName"]
	} else if type_ == "easy" {
		obj.RepositoryAddress = confVariable["fastapi"]["easyAddress"]
		obj.RepositoryName = confVariable["fastapi"]["easyName"]
	}
	fmt.Println(obj)
	obj.Generate()
}
