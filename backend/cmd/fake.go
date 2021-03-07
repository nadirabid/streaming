package cmd

import (
	"fmt"
	"os/exec"
	"streaming/utils"

	"github.com/spf13/cobra"
)

func fake(cmd *cobra.Command, args []string) error {
	// init viper
	// confFile, err := cmd.Flags().GetString("conf")
	// v, err := utils.ConfigureViper(confFile)
	// if err != nil {
	// 	return err
	// }

	// setup db

	// db, err := models.NewDatabase(v)
	// if err != nil {
	// 	return err
	// }

	// lets go: download & update db
	urls := map[string]string{
		"summer_adrif1": "https://www.youtube.com/watch?v=W1yaBmL42tY",
		"summer_adrif2": "https://www.youtube.com/watch?v=SuEZ3dJ4jxc",
		"summer_adrif3": "https://www.youtube.com/watch?v=-j6eKmu6qLE",
	}

	downloadSh := utils.GetAbsolutePath("assets/hls/download.sh")

	for name, url := range urls {
		cmdStr := fmt.Sprintf("%s -n %s -u '%s'", downloadSh, name, url)
		output, err := exec.Command("/bin/sh", cmdStr).Output()
		warner.Println(cmdStr)

		if err != nil {
			errer.Println(output, err.Error())
			return err
		}

		successer.Println(output)
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
