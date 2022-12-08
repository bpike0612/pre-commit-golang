package cmd

import (
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var in, output, version string

// rootCmd represents the base command when called without any subcommands.
var gitchglog = &cobra.Command{ //nolint
	Use:   "git-chglog [options] <tag query>",
	Short: "CHANGELOG generator implemented in Go (Golang). Anytime, anywhere, Write your CHANGELOG.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		var cmds []string

		if len(output) != 0 {
			cmds = append(cmds, "--output")
			cmds = append(cmds, "./CHANGELOG.md")
		}

		runTool("git-chglog", cmds)
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	gitchglog.PersistentFlags().StringVarP(&output, "output", "o", "", "output path and filename for the changelogs; if not specified, output to stdout")
	rootCmd.AddCommand(gitchglog)
}

func runTool(cmd string, args []string) {
	c := exec.Command(cmd, args...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	if err := c.Run(); err != nil {
		log.Fatalln(err)
	}
}
