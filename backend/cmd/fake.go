package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"streaming/models"
	"streaming/utils"
	"strings"

	"github.com/spf13/cobra"
	"gorm.io/gorm"
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

	basePath := fmt.Sprintf("assets/content/%s", "summer_adrift")
	fullBasePath := fmt.Sprintf("%s/%s", utils.GetBasePath(), basePath)

	content := &models.Content{}
	content.Name = "A summer adrift"
	content.Description = "Folks"
	content.AssetPath = basePath

	// check if it already exists
	err = db.Where("name = ?", content.Name).First(content).Error
	if err == nil {
		return errors.New("Already exists")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// lets go: download & update db
	urls := map[string]string{
		"One":   "https://www.youtube.com/watch?v=W1yaBmL42tY",
		"Two":   "https://www.youtube.com/watch?v=SuEZ3dJ4jxc",
		"Three": "https://www.youtube.com/watch?v=-j6eKmu6qLE",
	}

	downloadSh := utils.GetAbsolutePath("scripts/download.sh")
	series := []models.MiniSeries{}

	for name, url := range urls {
		successer.Println("Starting download: ", name, url)
		cmdStr := fmt.Sprintf("%s -f %s -n %s -u '%s'", downloadSh, fullBasePath, strings.ToLower(name), url)
		_, err := exec.Command("/bin/sh", "-c", cmdStr).Output()

		if err != nil {
			errer.Println("Failed to download: ", name, url)
			return err
		}

		successer.Println("Downloaded and processed HLS and thumbnails under: ", name, url)

		series = append(series, models.MiniSeries{
			MediaMetadata: models.MediaMetadata{
				Name:        name,
				Description: fmt.Sprintf("Episode %s", name),
				AssetPath:   fmt.Sprintf("%s/%s", basePath, strings.ToLower(name)),
			},
		})
	}

	content.MiniSeries = series

	if err := db.Create(content).Error; err != nil {
		return err
	}

	return nil
}

var fakeCmd = &cobra.Command{
	Use:   "fake",
	Short: "Fake data!",
	RunE:  fake,
}

func init() {
	dbCmd.AddCommand(fakeCmd)
}
