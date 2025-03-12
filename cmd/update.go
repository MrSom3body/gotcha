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
		return Release{}, fmt.Errorf("Request errored out: %s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Release{}, fmt.Errorf("Could not read response: %s", err)
	}

	var release Release
	err = json.Unmarshal(body, &release)
	if err != nil {
		return Release{}, fmt.Errorf("Cloud not parse response: %s", err)
	}

	return release, nil
}

func onNewestVersion(release Release) bool {
	if release.TagName == version {
		return true
	} else {
		return false
	}
}

func update(release Release) error {
	filePath := os.ExpandEnv("$HOME") + "/.local/bin/gotcha"
	resp, err := http.Get(release.Assets[0].BrowserDownloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	cmd := exec.Command("chmod", "+x", filePath)
	err = cmd.Run()
	if err != nil {
		return err
	}

	_, err = io.Copy(out, resp.Body)
	return err
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gotcha",
	Long:  "Update gotcha if there is a new version available",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checking for a new release...")
		release, err := getNewestRelease()
		if err != nil {
			log.Fatal(err)
		}

		if onNewestVersion(release) {
			fmt.Println("You're already on the newest version 🥳")
		} else {
			fmt.Println("Found a new release, updating...")
			err = update(release)
			if err != nil {
				log.Fatalf("Updating failed: %s", err)
			}
			fmt.Printf("You now have gotcha %s 🥳!\n", release.TagName)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
}
