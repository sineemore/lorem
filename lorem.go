package lorem

import (
	"bytes"
	"errors"
	"io"
)

// Holds lorem ipsum text. You can change it before actual usage
// to use other placeholder string.
var Lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."

// Convenient error to use within tests.
var ErrNotALorem = errors.New("Not a lorem")

type loremReader struct {
	pos int
}

func (r *loremReader) Read(p []byte) (readed int, err error) {

	for {
		n := copy(p[readed:], Lorem[r.pos:])
		if n == 0 {
			break
		}

		readed += n
		r.pos = (r.pos + n) % len(Lorem)
	}

	return
}

// NewLorem returns infinite Lorem reader that never
// returns io.EOF or any other error. Upon reaching
// end of Lorem reading will continue from beggining.
//
// Note: multiple Lorem are concated without space between them.
func NewLorem() io.Reader {
	return &loremReader{}
}

// NewLoremN reads n bytes of Lorem sequence and returns EOF.
func NewLoremN(n int) io.Reader {
	return io.LimitReader(NewLorem(), int64(n))
}

// Tests if provided buf holds a valid sequence of Lorem.
func IsLorem(buf []byte) bool {
	return IsLoremReader(bytes.NewReader(buf))
}

// IsLoremReader reads from r until io.EOF and tests if data
// is a valid sequence of Lorems. On error false is returned.
func IsLoremReader(r io.Reader) bool {
	var err error
	buf := make([]byte, len(Lorem))
	pos := 0
	for {
		var n int
		n, err = r.Read(buf)
		b := buf[:n]

		for len(b) > 0 {

			length := len(Lorem) - pos
			if length > len(b) {
				length = len(b)
			}

			if string(b[:length]) != Lorem[pos:pos+length] {
				return false
			}

			pos = (pos + length) % len(Lorem)
			b = b[length:]
		}

		if err != nil {
			break
		}
	}

	return err == io.EOF
}
