package cmd

import (
	"count-jobs/utils"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.Logger.Error(err)
		os.Exit(1)
	}
}
