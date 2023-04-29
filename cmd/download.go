package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download a file from the internet",
    Long: `Download a file from the internet long description`,
    Run: func(cmd *cobra.Command, args []string) {

        allFlag, _ := cmd.Flags().GetBool("all")
        urlFlag, _ := cmd.Flags().GetString("url")
        timeoutFlag, _ := cmd.Flags().GetInt("timeout")

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

        find_audios(u, allFlag, timeoutFlag)
    },
}

func init() {
    rootCmd.AddCommand(downloadCmd)
    downloadCmd.Flags().BoolP("all", "a", false, "Download all files in a book")
    downloadCmd.Flags().StringP("url", "u", "", "URL to download")
    downloadCmd.Flags().IntP("timeout", "t", 0, "HTTP timeout in seconds")
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

func find_audios(url *url.URL, all bool, timeout int) {
    fmt.Println("URL:", url)
    fmt.Println("Downloading file...")

    client := http.Client{}
    if timeout != 0 && timeout > 0 {
        client.Timeout = time.Duration(timeout) * time.Second
    }

    testUrl := "https://goldenaudiobook.com/rich-dad-poor-dad/"
    resp, err := client.Get(testUrl)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()
    fmt.Println("Status code:", resp.StatusCode)

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Finding media elements...")

    var audioUrls []string

    findMediaElements(&audioUrls, doc)

    var wg sync.WaitGroup
    wg.Add(len(audioUrls))

    for _, url := range audioUrls {
        go downloadAudio(url, &wg)
    }
    wg.Wait()
}

func findMediaElements(urls *[]string, doc *goquery.Document) {
    doc.Find(".wp-audio-shortcode").Each(func(i int, s *goquery.Selection) {
        *urls = append(*urls, s.Text())
    })
}

func downloadAudio(url string, wg *sync.WaitGroup) {
    defer wg.Done()

    fmt.Println("Downloading ", url)
    fileName := path.Base(url)
    output, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error while creating", fileName, "-", err)
        return
    }
    defer output.Close()

    response, err := http.Get(url)
    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    io.Copy(output, response.Body)
    fmt.Println("Downloaded", fileName)
}

