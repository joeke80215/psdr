package iountil

import "io"

//
//copy package stream to server
//
func Copy (w io.WriteCloser,r io.Reader) {
	defer w.Close()
	io.Copy(w,r)
}
