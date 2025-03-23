package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

type Release struct {
	TagName string `json:"tag_name"`
	Assets  []struct {
		Name               string `json:"name"`
		BrowserDownloadURL string `json:"browser_download_url"`
	} `json:"assets"`
}

const gotchaBinary = "$HOME/.local/bin/gotcha"

func fetchLatestRelease() (Release, error) {
	resp, err := http.Get("https://api.github.com/repos/MrSom3body/gotcha/releases/latest")
	if err != nil {
		return Release{}, fmt.Errorf("failed to fetch release info: %w", err)
	}
	defer resp.Body.Close()

	var release Release
	err = json.NewDecoder(resp.Body).Decode(&release)
	if err != nil {
		return Release{}, fmt.Errorf("failed to parse release data: %w", err)
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
		return fmt.Errorf("failed to download update: %w", err)
	}
	defer resp.Body.Close()

	out, err := os.Create(filePath + "~")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write update: %w", err)
	}

	_, err = os.Stat(filePath)
	if err == nil {
		err = os.Remove(filePath)
		if err != nil {
			return fmt.Errorf("failed to delete old version: %w", err)
		}
	}

	err = os.Rename(filePath+"~", filePath)
	if err != nil {
		return fmt.Errorf("failed to install new version: %w", err)
	}

	err = exec.Command("chmod", "+x", filePath).Run()
	if err != nil {
		return fmt.Errorf("failed to make gotcha executable: %w", err)
	}

	return nil
}

func installGotchaPrompt(filePath string) bool {
	fmt.Printf("Gotcha is not installed under %s. Do you want to install it there? [y/N] ", filePath)
	var input string
	fmt.Scanln(&input)
	return strings.ToLower(input) == "y"
}

func deleteOldGotcha(oldPath string) error {
	fmt.Print("Do you also want to delete the old gotcha? [Y/n] ")
	var input string
	fmt.Scanln(&input)
	if input != "" && strings.ToLower(input) != "y" {
		fmt.Printf("Old version not deleted. It remains at %s.\n", oldPath)
		return nil
	}

	info, err := os.Stat(oldPath)
	if err != nil {
		return fmt.Errorf("failed to fetch file info: %w", err)
	}

	if info.Mode()&os.ModeSymlink != 0 {
		oldPath, err = os.Readlink(oldPath)
		if err != nil {
			return fmt.Errorf("failed to resolve symlink: %w", err)
		}
	}

	if info.Mode().Perm()&0200 == 0 {
		return errors.New(fmt.Sprintf("can not delete read-only binary: %s", oldPath))
	}

	if err := os.Remove(oldPath); err != nil {
		return fmt.Errorf("failed to delete old gotcha: %w", err)
	}

	fmt.Println("Deleted old gotcha!")
	return nil
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update gotcha",
	Long:  "Update gotcha if there is a new version available",
	Run: func(cmd *cobra.Command, args []string) {
		filePath := os.ExpandEnv(gotchaBinary)

		gotchaPath, err := os.Executable()
		if err != nil {
			log.Fatalf("Failed to determine executable path: %v", err)
		}

		if gotchaPath != filePath {
			if installGotchaPrompt(filePath) {
				if err := deleteOldGotcha(gotchaPath); err != nil {
					log.Fatalf("Deleting old gotcha failed: %v", err)
				}
			} else {
				fmt.Println("Did not install gotcha.")
				return
			}
		}

		fmt.Println("Checking for a new release...")
		release, err := fetchLatestRelease()
		if err != nil {
			log.Fatal(err)
		}

		if isLatestVersion(release) {
			fmt.Println("You're already on the latest version 🥳")
			return
		}

		fmt.Println("New release found, updating...")
		if err := updateBinary(release, filePath); err != nil {
			log.Fatalf("Update failed: %v", err)
		}

		fmt.Printf("Update successful! Gotcha is now at version %s 🥳\n", release.TagName)
	}}

func init() {
	rootCmd.AddCommand(updateCmd)
}
