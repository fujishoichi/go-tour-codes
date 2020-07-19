package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}

func (rot13Reader rot13Reader) Read(b []byte) (int, error) {
	rn, re := rot13Reader.r.Read(b)
	if re != nil {
		return rn, re
	}
	for i, be := range b {
		b[i] = rot13(be)
	}
	return rn, nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	//s := strings.NewReader("Why did the chicken cross the road?")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
