//go:build amd64

package json

import (
	"io"

	"github.com/bytedance/sonic/encoder"
)

// Encode encodes the provided value into JSON format and writes it to the provided
// io.Writer.
//
// The function returns an error if there was an issue encoding the value.
func Encode(w io.Writer, v interface{}) error {
	// Create a new stream encoder using the provided io.Writer.
	enc := encoder.NewStreamEncoder(w)

	// Encode the provided value into the stream encoder.
	return enc.Encode(v)
}
