package base91

import (
	"io"

	"ekyu.moe/util/number"
)

func decode(src []byte) []byte {
	// b and n must be at least uint32
	var b, n uint = 0, 0
	v := -1
	decoded := []byte{}

	for i := 0; i < len(src); i++ {
		p, ok := dectab[src[i]]
		if !ok {
			// skip invalid character silently
			continue
		}
		if v < 0 {
			v = int(p)
			continue
		}
		v += int(p) * 91
		b |= uint(v) << n
		if v&0x1fff > 88 {
			n += 13
		} else {
			n += 14
		}
		for {
			decoded = append(decoded, uint8(b))
			b >>= 8
			n -= 8
			if n <= 7 {
				break
			}
		}
		v = -1
	}

	if v > -1 {
		decoded = append(decoded, uint8(b|uint(v)<<n))
	}

	return decoded
}

// Decode decodes src into dst. It returns the number of bytes written to dst.
func Decode(dst, src []byte) int {
	return copy(dst, decode(src))
}

// DecodeString returns the bytes represented by the base91 string s, probably what
// you want.
func DecodeString(s string) []byte {
	return decode([]byte(s))
}

type decoder struct {
	reader io.Reader
	buf    []byte

	b, n uint
	v    int
}

// NewDecoder constructs a new base91 stream decoder. Data read from the returned
// reader is base91 decoded from r.
func NewDecoder(r io.Reader) io.Reader {
	return &decoder{
		reader: r,
		buf:    []byte{},
		b:      0,
		n:      0,
		v:      -1,
	}
}

func (d *decoder) read(c []byte) (int, error) {
	// 只是个估计，base91 的转码率是不固定的
	encodedLenEst := int(number.Round(1.23078*float64(len(c))+0.36812, 0))
	encodedBuf := make([]byte, encodedLenEst)

	upstreamReadLen, err := d.reader.Read(encodedBuf)
	if err != nil {
		return 0, err
	}

	for i := 0; i < upstreamReadLen; i++ {
		p, ok := dectab[encodedBuf[i]]
		if !ok {
			continue
		}
		if d.v < 0 {
			d.v = int(p)
			continue
		}
		d.v += int(p) * 91
		d.b |= uint(d.v) << d.n
		if d.v&0x1fff > 88 {
			d.n += 13
		} else {
			d.n += 14
		}
		for {
			d.buf = append(d.buf, uint8(d.b))
			d.b >>= 8
			d.n -= 8
			if d.n <= 7 {
				break
			}
		}
		d.v = -1
	}

	if d.v > -1 {
		d.buf = append(d.buf, uint8(d.b|uint(d.v)<<d.n))
	}

	n := copy(c, d.buf)
	if n < len(c) {
		m, err := d.read(c[n:])
		return m + n, err
	}

	return n, nil
}

func (d *decoder) Read(c []byte) (int, error) {
	// flush buffer if there is any
	n := copy(c, d.buf)
	d.buf = d.buf[:0]
	if n > 0 {
		m, err := d.read(c[n:])
		return m + n, err
	}

	return d.read(c)
}
