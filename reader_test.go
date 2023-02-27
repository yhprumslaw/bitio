package bitio

import (
	"bytes"
	"testing"
)

func TestNewBitReader(t *testing.T) {
	br := NewReader(
		bytes.NewReader([]byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}))
	testCases := []struct {
		n   uint8
		exp uint64
	}{
		{0, 0},
		{1, 1},
		{2, 3},
		{3, 7},
		{4, 15},
		{5, 31},
		{6, 63},
		{7, 127},
		{8, 255},
	}

	for _, tc := range testCases {
		if got, _ := br.ReadBits(tc.n); got != tc.exp {
			t.Errorf("ReadBits(%d) = %d, want %d", tc.n, got, tc.exp)
		}
	}
}
