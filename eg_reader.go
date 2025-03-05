package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rr rot13Reader) Read(p []byte) (int, error) {
	c, err := rr.r.Read(p)
	if err == nil {
		for i, c := range p {
			p[i] = rot13(c)
		}
	}
	return c, err
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
