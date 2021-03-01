package cmd

import (
	"streaming/models"
	"streaming/utils"

	"github.com/spf13/cobra"
)

func fake(cmd *cobra.Command, args []string) error {
	// init viper
	confFile, err := cmd.Flags().GetString("conf")
	v, err := utils.ConfigureViper(confFile)
	if err != nil {
		return err
	}

	// setup db

	db, err := models.NewDatabase(v)
	if err != nil {
		return err
	}

	for i := 1; i <= 3; i++ {
		video := &models.Video{
			Name:        "Part 1",
			Description: "Starting the journey",
		}
	}

	return nil
}

var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Fake data!",
	RunE:  fake,
}

func init() {
	rootCmd.AddCommand(fakeCmd)
}
