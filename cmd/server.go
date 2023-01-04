package cmd

import (
	"count-jobs/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := api.New().Run()
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
