package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{ //nolint
	Use:   "git-chglog --output CHANGELOG.md",
	Short: "git-chglog generates changelog",
	Long:  `git-chglog generates changelog`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
	run()
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.studybuddy.yaml)")
	var flag string
	// rootCmd.PersistentFlags().StringVar(&flag, "output", "CHANGELOG.md", "")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&flag,
		"output",
		"o",
		"",
		"output path and filename for the changelogs; if not specified, output to stdout")
}

func run() {
	out, err := exec.Command("git-chglog", "--output CHANGELOG.md").Output()
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println("Command Successfully Executed")
	output := string(out[:])
	fmt.Println(output)
}
