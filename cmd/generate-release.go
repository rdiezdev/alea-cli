package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/exp/slices"
)

var basePath = viper.GetString("projects_folder")

func init() {
	rootCmd.AddCommand(generateReleaseCmd)
}

var generateReleaseCmd = &cobra.Command{
	Use:   "release",
	Short: "Generate a release for a project of your choice",
	Long:  `Generate a release for a project of your choice`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		availableProjects := getProjectFolderNames()

		var projectSelected string

		if len(args) == 0 {
			prompt := promptui.Select{
				Label:             "Select a project",
				Items:             availableProjects,
				StartInSearchMode: true,
				Searcher: func(input string, index int) bool {
					item := availableProjects[index]
					name := strings.Replace(strings.ToLower(item), " ", "", -1)
					input = strings.Replace(strings.ToLower(input), " ", "", -1)
					return strings.Contains(name, input)
				},
			}
			_, projectSelected, _ = prompt.Run()
		} else {

			if !slices.Contains(availableProjects, args[0]) {
				fmt.Println("Project not found")
				os.Exit(1)
			}

			projectSelected = args[0]
		}

		releaseVersion := getReleaseVersion(projectSelected)

		releaseBranchName := "release/v" + releaseVersion

		createReleaseAndPushIt(releaseBranchName, projectSelected)
	},
}

func getProjectFolderNames() (projectFolderNames []string) {
	dirInfo, err := ioutil.ReadDir(basePath)

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	for _, f := range dirInfo {
		if f.IsDir() {
			// is a git repo?

			if _, err := os.Stat(basePath + f.Name() + "/.git"); err == nil {
				projectFolderNames = append(projectFolderNames, f.Name())
			}
		}
	}

	sort.Slice(projectFolderNames, func(i, j int) bool {
		return projectFolderNames[i] < projectFolderNames[j]
	})

	return
}

func createReleaseAndPushIt(releaseBranchName string, dirName string) {
	os.Chdir("/Users/ramonespana/alea/bitbucket/" + dirName)

	cmd := exec.Command("git", "checkout", "develop")

	if _, err := cmd.Output(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "checkout", "-b", releaseBranchName)

	if _, err := cmd.Output(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "push", "origin", releaseBranchName)

	if _, err := cmd.Output(); err != nil {
		log.Fatal(err)
	}
}

func getReleaseVersion(dirName string) string {
	os.Chdir(basePath + dirName)
	cmd := exec.Command("mvn", "help:evaluate", "-Dexpression=project.version", "-q", "-DforceStdout")

	out, err := cmd.Output()

	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(out), "-")[0]
}
