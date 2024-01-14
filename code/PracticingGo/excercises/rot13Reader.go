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
	switch {
	case 'A' <= b && b <= 'Z':
		return (b-'A'+13)%26 + 'A'
	case 'a' <= b && b <= 'z':
		return (b-'a'+13)%26 + 'a'
	default:
		return b
	}
}

func (p *rot13Reader) Read(output []byte) (int, error) {
	for {
		n, err := p.r.Read(output)
		if err != nil {
			return 0, err
		}
		for index, item := range output {
			if (item >= 'a' && item <= 'z') || (item >= 'A' && item <= 'Z') {
				output[index] = rot13(item)
			}
		}
		// below return is important so that function
		// calling this read can recognize that we
		// have wrote something in output
		return n, nil
	}
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
