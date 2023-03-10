package main

import (
	"log"

	archive "github.com/nikita/assign3/archive"
	download "github.com/nikita/assign3/dowload"
)

func main() {

	fullURLFile := "https://filesamples.com/samples/video/mp4/sample_640x360.mp4"
	src := "/Users/nikita.mogha/Downloads/sample_960x400_ocean_with_audio.mp4"

	var myweb download.Downloader
	myweb = download.NewDownloader()
	r1, err := myweb.Download(fullURLFile)

	if err != nil {
		log.Fatal(err)
	}

	myweb = download.NewDownloaderFromFileSystem()

	r2, err := myweb.Download(src)

	if err != nil {
		log.Fatal(err)
	}

	var zipper archive.Archiver
	zipper = archive.NewArchiver()

	_, err = zipper.Archive([]string{"f1", "f2"}, r1, r2)

	if err != nil {
		log.Fatal(err)
	}

}
