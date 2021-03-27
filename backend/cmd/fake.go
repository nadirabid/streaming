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

	// content := &db.Content{
	// 	Name:        "A summer adrift",
	// 	Description: "Folks",
	// }

	// lets go: download & update db
	urls := map[string]string{
		"One":   "https://www.youtube.com/watch?v=W1yaBmL42tY",
		"Two":   "https://www.youtube.com/watch?v=SuEZ3dJ4jxc",
		"Three": "https://www.youtube.com/watch?v=-j6eKmu6qLE",
	}

	downloadSh := utils.GetAbsolutePath("assets/hls/download.sh")

	fmt.Println("hello2")

	for name, url := range urls {
		cmdStr := fmt.Sprintf("%s -f %s -n %s -u '%s'", downloadSh, "summer_adrift", name, url)
		fmt.Println(cmdStr)
		_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()

		fmt.Println(downloadSh, "hello")

		if err != nil {
			fmt.Println("errr")
			return err
		}

		successer.Println("Downloaded: ", name, url)
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
