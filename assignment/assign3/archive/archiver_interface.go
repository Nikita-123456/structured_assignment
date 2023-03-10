package archive

import "io"

type Archiver interface {
	Archive([]string, ...io.Reader) (io.Reader , error)
}


