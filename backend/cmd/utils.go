package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"streaming/utils"
)

func configureViperFromCmd(cmd *cobra.Command) (*viper.Viper, error) {
	confFile, err := cmd.Flags().GetString("conf")
	if err != nil {
		return nil, err
	}

	return utils.ConfigureViper(confFile)
}
