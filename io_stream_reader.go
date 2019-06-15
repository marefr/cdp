package cdp

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"strings"

	cdpio "github.com/mafredri/cdp/protocol/io"
)

// IOStreamReader represents a stream reader.
type IOStreamReader struct {
	next func(pos, size int) (*cdpio.ReadReply, error)
	r    io.Reader
	pos  int
	eof  bool
}

// NewIOStreamReader returns a new reader for io.Streams.
func NewIOStreamReader(ctx context.Context, c *Client, handle cdpio.StreamHandle) *IOStreamReader {
	args := cdpio.NewReadArgs(handle)
	return &IOStreamReader{
		next: func(pos, size int) (*cdpio.ReadReply, error) {
			args.SetOffset(pos).SetSize(size)
			return c.IO.Read(ctx, args)
		},
	}
}

func (r *IOStreamReader) read(p []byte) (n int, err error) {
	n, err = r.r.Read(p)
	r.pos += n
	if !r.eof && err == io.EOF {
		r.r = nil
		err = nil
	}
	return n, err
}

func (r *IOStreamReader) Read(p []byte) (n int, err error) {
	if r.r != nil {
		// Continue reading from buffer.
		return r.read(p)
	}
	if r.eof {
		return 0, io.EOF
	}
	if len(p) == 0 {
		return 0, nil
	}

	// Chrome might have an off-by-one when deciding the maximum
	// size (at least for base64 encoded data), usually it will
	// overflow. We subtract one to make sure it fits into p.
	size := len(p) - 1
	if size < 1 {
		// Safety-check to avoid crashing Chrome (e.g. via SetSize(-1)).
		size = 1
	}

	reply, err := r.next(r.pos, size)
	if err != nil {
		return 0, err
	}

	r.eof = reply.EOF

	switch {
	case reply.Base64Encoded != nil && *reply.Base64Encoded:
		b := []byte(reply.Data)
		size := base64.StdEncoding.DecodedLen(len(b))

		// Safety-check for fast-path to avoid panics.
		if len(p) >= size {
			n, err = base64.StdEncoding.Decode(p, b)
			r.pos += n
			return n, err
		}

		r.r = base64.NewDecoder(base64.StdEncoding, bytes.NewReader(b))
	default:
		r.r = strings.NewReader(reply.Data)
	}

	return r.read(p)
}
