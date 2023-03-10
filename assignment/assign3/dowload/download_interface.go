package download

import "io"

type Downloader interface {
	Download(string) (io.Reader, error)
}
