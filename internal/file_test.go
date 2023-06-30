package internal_test

import (
	"io"
	"testing"

	"github.com/hirosuke29/fcv/internal"
)

func TestDecode(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		inFile       io.ReadWriteCloser
		inFCVObjects []internal.FCVObject
		expected     []byte
	}{
		{
			name: "test",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
