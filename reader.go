package bitio

import (
	"errors"
	"io"
)

var (
	ExceededErr = errors.New("parameter out of range: cannot be greater than 64")
)

type Reader struct {
	r io.Reader

	n         uint8 // left for bitsCache, max is 7
	bitsCache uint64
	readCache [8]byte

	count uint64 // bits count has read
	err   error
}

// NewReader create a new Reader
func NewReader(r io.Reader) *Reader {
	return &Reader{
		r: r,
	}
}

func (r *Reader) ReadBits(n uint8) (bits uint64, err error) {
	if n > 64 {
		return 0, ExceededErr
	}
	if n > r.n {
		bits = r.bitsCache << (n - r.n)
	}

	want := (n - r.n + 7) / 8
	_, err = io.ReadFull(r.r, r.readCache[:want])
	if err != nil {
		return 0, err
	}
	for _, i := range r.readCache[:want] {
		r.bitsCache <<= 8
		r.bitsCache |= uint64(i)
	}
	r.n += want * 8

	bits |= r.bitsCache >> uint(r.n-n)
	r.bitsCache ^= bits << uint(r.n-n)
	r.n -= n
	r.count += uint64(n)

	return bits, nil

}

func (r *Reader) ReadCount() uint64 {
	return r.count
}
