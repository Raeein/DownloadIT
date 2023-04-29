package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands

var rootCmd = &cobra.Command{
    Use:   "DownloadIT",
    Short: "DownloadIT is a CLI tool for downloading files from the internet",
    Long: `DownloadIT is a CLI tool for downloading files from the internet.
    It supports downloading files from HTTP, HTTPS, and FTP protocols.
    It also supports downloading files from Google Drive and Dropbox.
    `,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}


