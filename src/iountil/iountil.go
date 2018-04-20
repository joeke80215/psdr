package iountil

import "io"

func Copy (w io.WriteCloser,r io.Reader) {
	defer w.Close()
	io.Copy(w,r)
}
