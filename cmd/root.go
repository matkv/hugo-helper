package cmd

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/matkv/utils/internal/config"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "utils",
	Short: "CLI tool to automate some personal tasks",
	Long:  `CLI tool to automate some personal tasks. Some of the tasks include managing dotfiles, managing my hugo website and updating my workout tracker.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Initializing configuration...")
	cobra.OnInitialize(config.InitConfig)
}
