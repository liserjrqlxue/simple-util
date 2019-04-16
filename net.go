package simple_util

import (
	"io"
	"net/http"
	"os"
)

// https://golangcode.com/download-a-file-from-a-url/
// DownloadFile will download a url to a local file.
// It's efficient because it will write as it downloads and not load the whole file into memroy.
func DownloadFile(filepath string, url string) error {
	// Get the data
	resp, err := http.Get(url)
	CheckErr(err)
	defer resp.Body.Close()

	// Create teh file
	out, err := os.Create(filepath)
	CheckErr(err)
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
