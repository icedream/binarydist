package binarydist

import (
	"io"

	"github.com/dsnet/compress/bzip2"
)

func newBzip2Writer(w io.Writer) (wc io.WriteCloser, err error) {
	wc, err = bzip2.NewWriter(w, nil)
	return
}
