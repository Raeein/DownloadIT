package cmd

import (
	"fmt"
	"net/url"

	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download a file from the internet",
    Long: `Download a file from the internet long description`,
    Run: func(cmd *cobra.Command, args []string) {

        allFlag, _ := cmd.Flags().GetBool("all")

        urlFlag, _ := cmd.Flags().GetString("url")

        if urlFlag == "" {
            fmt.Println("URL is required")
            cmd.Help()
            return
        }

        u, err := url.ParseRequestURI(urlFlag)
        if err != nil {
            fmt.Println("Invalid URL")
            cmd.Help()
            return
        }

        downloadFile(u, allFlag)
    },
}

func init() {
    rootCmd.AddCommand(downloadCmd)
    downloadCmd.Flags().BoolP("all", "a", false, "Download all files in a book")
    downloadCmd.Flags().StringP("url", "u", "", "URL to download")
}

// each piece is a part of a music file to be download and concatenated laters on
type piece struct {
    url string
    start int
    end int
    order int
    downloaded bool
}

type book []piece

func (b *book) addPiece(newPiece piece) {
    *b = append(*b, newPiece)
}


func downloadFile(url *url.URL, all bool) {
    fmt.Println("URL:", url)
    fmt.Println("Downloading file...")
}

