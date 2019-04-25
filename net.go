package simple_util

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"net/http"
	"os"
	"strings"
)

// From https://golangcode.com/download-a-file-from-a-url/
// DownloadFile will download a url to a local file.
// It's efficient because it will write as it downloads and not load the whole file into memroy.
func DownloadFile(filepath, url string) error {
	// Get the data
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	CheckErr(err)
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

// From https://golangcode.com/download-a-file-with-progress/
// WriterCounter counts the number of bytes written to it.
// It implements to the io.Writer interface and we can pass this into io.TeeReader() which will report progress on each write cycle.
type WriterCounter struct {
	Total uint64
}

func (wc *WriterCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

func (wc *WriterCounter) PrintProgress() {
	// Clear the line by using a character return to go back to teh start an remove the remaining characters by filing it with spaces
	fmt.Printf("\t%", strings.Repeat(" ", 35))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10MB)
	fmt.Println("\rDownloading.. %s complete", humanize.Bytes(wc.Total))
}

// We pass an io.TeeReader into Copy() to report progress on the download.
func DownloadFileProgress(filepath, url string) {
	// Create the file,but give it a tmp file extension,
	// this means we won't overwrite a file until it's downloaded,
	// but we'll remove the tmp extension once download.
	out, err := os.Create(filepath + ".tmp")
	CheckErr(err)
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	println(resp.ContentLength)
	// Create out progress reporter and pass it to be ued alongside our writer
	counter := &WriterCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	CheckErr(err)

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Print("\n")

	err = os.Rename(filepath+".tmp", filepath)
	CheckErr(err)
}
