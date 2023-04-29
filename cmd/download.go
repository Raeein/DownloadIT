package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

type options struct {
    all bool
    url *url.URL
    timeout int
    merge bool
    delete bool
    verbose bool
}

var downloadCmd = &cobra.Command{
    Use:   "download",
    Short: "Download audio books",
    Long: `Download audio books from https://goldenaudiobook.com/ for free`,
    Run: func(cmd *cobra.Command, args []string) {

        o := options{}

        urlFlag, _ := cmd.Flags().GetString("url")
        allFlag, _ := cmd.Flags().GetBool("all")
        timeoutFlag, _ := cmd.Flags().GetInt("timeout")
        mergeFlag, _ := cmd.Flags().GetBool("merge")
        deleteFlag, _ := cmd.Flags().GetBool("delete")
        verboseFlag, _ := cmd.Flags().GetBool("verbose")

        o.all = allFlag
        o.timeout = timeoutFlag
        o.merge = mergeFlag
        o.delete = deleteFlag
        o.verbose = verboseFlag

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

        o.url = u

        find_audios(o)
    },
}

func init() {
    rootCmd.AddCommand(downloadCmd)
    downloadCmd.Flags().BoolP("all", "a", true, "Download all the audio files - Default is true")
    downloadCmd.Flags().StringP("url", "u", "", "URL to download audio books")
    downloadCmd.Flags().IntP("timeout", "t", 0, "HTTP timeout in seconds - Default is 0")
    downloadCmd.Flags().BoolP("merge", "m", false, "Merge all the audio files into one - Default is false")
    downloadCmd.Flags().BoolP("delete", "d", false, "Delete temp files after downloading - Default is false")
    downloadCmd.Flags().BoolP("verbose", "v", false, "Show output of ffmpeg command - Default is false")
}

func find_audios(o options) {

    client := http.Client{}
    if o.timeout != 0 && o.timeout > 0 {
        client.Timeout = time.Duration(o.timeout) * time.Second
    }

    resp, err := client.Get(o.url.String())
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    if status := resp.StatusCode; status != http.StatusOK {
        fmt.Println("Error: status code", status)
        return
    }

    doc, err := goquery.NewDocumentFromReader(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Finding audio files...")

    var audioUrls []string

    findMediaElements(&audioUrls, doc)

    var wg sync.WaitGroup
    wg.Add(len(audioUrls))

    for _, url := range audioUrls {
        go downloadAudio(url, &wg, o.timeout)
    }
    wg.Wait()

    fmt.Println("")
    fmt.Println("Finished downloading files")

    if !o.merge {
        return
    }
    var filenames []string
    for _, url := range audioUrls {
        filenames = append(filenames, path.Base(url))
    }

    outputName := strings.Split(path.Base(o.url.Path), ".")[0] + ".mp3"
    fmt.Println("Merging files...")
    mergeFiles(outputName, filenames, o.verbose)

    if !o.delete {
        return
    }

    fmt.Println("Deleting temp files...")
    for _, filename := range filenames {
        err := os.Remove(filename)
        if err != nil {
            fmt.Println("Error while deleting", filename, "-", err)
        }
    }
    fmt.Println("Finished deleting temp files")
    fmt.Println("Enjoy your educational research 0_o!")

}

func mergeFiles(outputName string, audioUrls []string, verbose bool) {
    _, err := exec.LookPath("ffmpeg")
    if err != nil {
        fmt.Println("ffmpeg not found in PATH")
        return
    }
	args := []string{"-y", "-i", "concat:" + strings.Join(audioUrls, "|"), "-acodec", "copy", outputName}
    cmd := exec.Command("ffmpeg", args...)
    if verbose {
        cmd.Stdout = os.Stdout
        cmd.Stderr = os.Stderr
    }

    err = cmd.Run()
    if err != nil {
        fmt.Println("Error while merging files:", err)
    }
    fmt.Println("Merging finished")
    fmt.Println("")
}

func findMediaElements(urls *[]string, doc *goquery.Document) {
    doc.Find(".wp-audio-shortcode").Each(func(i int, s *goquery.Selection) {
        *urls = append(*urls, s.Text())
    })
}

func downloadAudio(url string, wg *sync.WaitGroup, timeout int) {
    defer wg.Done()

    fmt.Println("Downloading ", url)
    fileName := path.Base(url)
    output, err := os.Create(fileName)
    if err != nil {
        fmt.Println("Error while creating", fileName, "-", err)
        return
    }
    defer output.Close()

    client := http.Client{}
    if timeout != 0 && timeout > 0 {
        client.Timeout = time.Duration(timeout) * time.Second
    }
    response, err := client.Get(url)

    if err != nil {
        fmt.Println("Error while downloading", url, "-", err)
        return
    }
    defer response.Body.Close()

    io.Copy(output, response.Body)
    fmt.Println("Downloaded", fileName)
}

