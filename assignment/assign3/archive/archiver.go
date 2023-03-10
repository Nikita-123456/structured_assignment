package archive

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
)

type Zipper struct {
}

func NewArchiver() *Zipper {
	return &Zipper{}
}

func (ar *Zipper) Archive(name []string, readers ...io.Reader) (io.Reader, error) {
	zip_ptr, err := os.Create("archive.zip")

	if err != nil {
		log.Fatal(err)
	}
	zip_writer := zip.NewWriter(zip_ptr)
	for i, r := range readers {
		file, err := zip_writer.Create(name[i])
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(file, r)
		if err != nil {
			log.Fatal(err)
		}

	}
	fmt.Println("Done with making the zip file")

	return zip_ptr, nil
}
