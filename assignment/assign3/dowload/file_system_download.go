package download

import (
	"fmt"
	"io"
	"log"
	"os"
)

type FileSystem struct {
}

func NewDownloaderFromFileSystem() *FileSystem {
	return &FileSystem{}
}

func (v2 *FileSystem) Download(path string) (io.Reader, error) {
	// copying content of source file into destination file

	source_file_ptr, err := os.Open(path)
	defer source_file_ptr.Close()
	if err != nil {
		log.Fatal(err)
	}
	dest_file_ptr, err := os.Create("copy_of_video.mp4")
	if err != nil {
		log.Fatal(err)
	}
	_, err = io.Copy(dest_file_ptr, source_file_ptr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Downloaded a file from file system")
	dest_file_ptr.Seek(0,io.SeekStart)
	return dest_file_ptr, nil

}
