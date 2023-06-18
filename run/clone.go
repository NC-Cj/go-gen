package run

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type CloneObject struct {
	RepositoryAddress string
	RepositoryName    string
	ProjectName       string
	Prisma            bool
}

var ee = "Encountered an error: "

func (c *CloneObject) Generate() {
	//fmt.Println("Get remote repository...")
	c.clone()
	c.removeGitFolder()
	c.renameProjectName()
	time.Sleep(time.Second * 2)

	if c.Prisma {
		checkNPM()
		checkPrisma()
		c.runPrisma()
	}

	fmt.Println(`OK, I'm done, Have fun programming :)
		"Go Enter "code ."`)
}

func (c *CloneObject) clone() {
	cmd := exec.Command("git", "clone", c.RepositoryAddress)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Please check if you can clone the github project normally")
		os.Exit(1)
	}
}

func (c *CloneObject) removeGitFolder() {
	pwd, _ := os.Getwd()
	gitDir := pwd + c.RepositoryName + "/.git"
	err := os.RemoveAll(gitDir)
	if err != nil {
		fmt.Println(ee, err)
		os.Exit(1)
	}
}

func (c *CloneObject) renameProjectName() {
	oldpath := "./" + c.RepositoryName
	err := os.Rename(oldpath, c.ProjectName)
	if err != nil {
		fmt.Println(ee, err)
		os.Exit(1)
	}
}

//func (c *CloneObject) checkPrismaPack() {
//	fmt.Println("Check execution environment...")
//	npm := exec.Command("npm", "list", "-g", "prisma")
//	output, err := npm.CombinedOutput()
//	if err != nil {
//		fmt.Println(ee, err)
//		os.Exit(1)
//	}
//
//	strOutput := string(output)
//	if find := strings.Contains(strOutput, "prisma"); find {
//		return
//	}
//	fmt.Println("Downloading Prisma, please make sure `npm` is installed on the computer...")
//	install := exec.Command("npm", "install", "-g", "prisma")
//	cmdRun(install)
//	fmt.Println("Successfully installed package...")
//}

func (c *CloneObject) runPrisma() {
	err := os.Chdir(".\\" + c.ProjectName + "\\")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	npx := exec.Command("npx", "prisma", "init")
	cmdRun(npx)
	fmt.Println("prisma Initialized successfully...")
}

func cmdRun(cmd *exec.Cmd) {
	if err := cmd.Run(); err != nil {
		fmt.Println(ee, err)
		os.Exit(1)
	}
}
func cmdOutput(cmd *exec.Cmd) string {
	output, _ := cmd.CombinedOutput()
	strOutput := string(output)
	return strOutput
}
func checkNPM() {
	cmd := exec.Command("npm", "-v")
	s := cmdOutput(cmd)
	if s == "" {
		fmt.Println("Missing necessary environment, please install `node` or `npm` package management")
		os.Exit(0)
	}
	return
}

func checkPrisma() {
	cmd := exec.Command("npm", "list", "-g", "prisma")
	s := cmdOutput(cmd)
	if ok := strings.Contains(s, "prisma"); ok {
		return
	}
	fmt.Println("Missing necessary environment, you need to execute `npm install - g prisma` separately")
	os.Exit(0)
}
