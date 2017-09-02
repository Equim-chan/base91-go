package base91

import (
	"io"
)

type wrapper struct {
	// output
	writer io.Writer

	buf     []byte
	lastPos int
}

// wrapEncoder serves as a bridge.
type wrapEncoder struct {
	encoder io.WriteCloser
	wrapper *wrapper
}

// NewLineWrapper returns an base91 encoeder that encode data and insert CRLF every 76
// characters. This is useful for emails.
func NewLineWrapper(w io.Writer) io.WriteCloser {
	wr := &wrapper{
		writer:  w,
		buf:     make([]byte, 78),
		lastPos: 0,
	}
	we := &wrapEncoder{
		// It proxies data written to encoder to wrapper.
		encoder: NewEncoder(wr),
		wrapper: wr,
	}

	return we
}

func (we *wrapEncoder) Write(p []byte) (int, error) {
	return we.encoder.Write(p)
}

func (we *wrapEncoder) Close() error {
	if err := we.encoder.Close(); err != nil {
		return err
	}

	return we.wrapper.Close()
}

func (wr *wrapper) flush() error {
	wr.buf[wr.lastPos] = '\r'
	wr.lastPos++
	wr.buf[wr.lastPos] = '\n'
	wr.lastPos++

	if _, err := wr.writer.Write(wr.buf[:wr.lastPos]); err != nil {
		return err
	}

	wr.lastPos = 0

	return nil
}

func (wr *wrapper) Write(p []byte) (int, error) {
	var i int
	for i = 0; i < len(p); i++ {
		if wr.lastPos == 76 {
			if err := wr.flush(); err != nil {
				return i + 1, err
			}
		}

		wr.buf[wr.lastPos] = p[i]
		wr.lastPos++
	}

	return i, nil
}

func (wr *wrapper) Close() error {
	return wr.flush()
}
