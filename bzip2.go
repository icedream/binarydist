package binarydist

import (
	"io"
	"os/exec"

	"github.com/dsnet/compress/bzip2"
)

type bzip2Writer struct {
	c *exec.Cmd
	w io.WriteCloser
}

func (w bzip2Writer) Write(b []byte) (int, error) {
	return w.w.Write(b)
}

func (w bzip2Writer) Close() error {
	if err := w.w.Close(); err != nil {
		return err
	}
	return w.c.Wait()
}

// Package compress/bzip2 implements only decompression,
// so we'll fake it by running bzip2 in another process.
func newBzip2Writer(w io.Writer) (wc io.WriteCloser, err error) {
	compressedWriter, err := bzip2.NewWriter(w, nil)
	if err != nil {
		return
	}

	wc = compressedWriter
	return
}
