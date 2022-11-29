package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var gitchglog = &cobra.Command{ //nolint
	Use:   "git-chglog [options] <tag query>",
	Short: "CHANGELOG generator implemented in Go (Golang). Anytime, anywhere, Write your CHANGELOG.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		cmds := []string{"--output CHANGELOG.md"}
		runTool("git-chglog", append(cmds, args...))
	},
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.studybuddy.yaml)")

	// var flag string

	// rootCmd.PersistentFlags().StringVar(&flag, "output", "CHANGELOG.md", "")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	//rootCmd.Flags().StringVarP(&flag,
	//	"output",
	//	"o",
	//	"",
	//	"output path and filename for the changelogs; if not specified, output to stdout")
	rootCmd.AddCommand(gitchglog)
}

func runTool(cmd string, args []string) {
	c := exec.Command(cmd, args...)

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr

	err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error running %s %s: %s", cmd, args, err)
		os.Exit(1)
	}
}
