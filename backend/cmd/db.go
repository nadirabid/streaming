package cmd

import (
	"github.com/spf13/cobra"

	"streaming/models"
)

func migrate(cmd *cobra.Command, args []string) error {
	// init viper
	v, err := configureViperFromCmd(cmd)
	if err != nil {
		return err
	}

	if err := models.Migrate(v); err != nil {
		return err
	}

	return nil
}
