package binarydist

import (
	"io"

	"github.com/dsnet/compress/bzip2"
)

func newBzip2Writer(w io.Writer) (wc io.WriteCloser, err error) {
	compressedWriter, err := bzip2.NewWriter(w, nil)
	if err != nil {
		return
	}

	wc = compressedWriter
	return
}
