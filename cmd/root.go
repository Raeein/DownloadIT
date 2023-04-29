package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "DownloadIT",
    Short: "DownloadIT is a CLI tool for downloading files from the internet",
    Long: `DownloadIT is a CLI tool for downloading free audio books from the internet.
    The only supported website is https://goldenaudiobook.com/
    `,
}

func Execute() {
    err := rootCmd.Execute()
    if err != nil {
        os.Exit(1)
    }
}


