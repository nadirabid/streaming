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

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Apply any new migrations",
	RunE:  migrate,
}

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Commands to interact with database",
}

func init() {
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(migrateCmd)
}
