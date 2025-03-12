package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

func getNewestRelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/MrSom3body/gotcha/releases/latest")
	if err != nil {
		return Release{}, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	var release Release
	if err := json.NewDecoder(resp.Body).Decode(&release); err != nil {
		return Release{}, fmt.Errorf("failed to parse response: %w", err)
	}
	return release, nil
}

func isLatestVersion(release Release) bool {
	return release.TagName == version
}

func getAssetURL(release Release, assetName string) (string, error) {
	for _, asset := range release.Assets {
		if asset.Name == assetName {
			return asset.BrowserDownloadURL, nil
		}
	}
	return "", fmt.Errorf("asset %s not found", assetName)
}

func updateBinary(release Release, filePath string) error {
	assetURL, err := getAssetURL(release, "gotcha")
	if err != nil {
		return err
	}

	resp, err := http.Get(assetURL)
	if err != nil {
		return fmt.Errorf("download failed: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("file creation failed: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, resp.Body); err != nil {
		return fmt.Errorf("file write failed: %w", err)
	}

	return exec.Command("chmod", "+x", filePath).Run()
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gotcha",
	Long:  "Update gotcha if there is a new version available",
	Run: func(cmd *cobra.Command, args []string) {
		filePath, err := os.Executable()
		if err != nil {
			log.Fatal("Failed to determine executable path:", err)
		}

		if fileInfo, err := os.Stat(filePath); err != nil || fileInfo.Mode()&0200 == 0 {
			log.Fatal("Gotcha is installed in a read-only location, can't update!")
		}

		fmt.Println("Checking for a new release...")
		release, err := getNewestRelease()
		if err != nil {
			log.Fatal(err)
		}

		if isLatestVersion(release) {
			fmt.Println("You're already on the newest version 🥳")
			return
		}

		fmt.Println("Found a new release, updating...")
		if err := updateBinary(release, filePath); err != nil {
			log.Fatalf("Updating failed: %v", err)
		}
		fmt.Printf("You now have gotcha %s 🥳!\n", release.TagName)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
