package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type UrlVideo struct {
}

func NewDownloader() *UrlVideo {
	return &UrlVideo{}
}
func (v *UrlVideo) Download(myurl string) (io.Reader, error) {
	// Create blank file

	fileURL, err := url.Parse(myurl)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	nameOfFile := segments[len(segments)-1]

	file, err := os.Create(nameOfFile)
	if err != nil {
		return nil, err
	}

	client := http.Client{}

	// Put content on file

	resp, err := client.Get(myurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Downloaded a file from url")
	file.Seek(0, io.SeekStart)
	return file, nil
}
