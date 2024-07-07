//go:build !amd64

package json

import (
	"io"

	goccy "github.com/goccy/go-json"
)

// Encode encodes the provided value into JSON format and writes it to the provided writer.
// It returns an error if any error occurs during the encoding process.
//
// Parameters:
// - w: The writer to which the JSON encoded value will be written.
// - v: The value to be encoded into JSON format.
//
// Returns:
// - error: An error if any error occurs during the encoding process, otherwise nil.
func Encode(w io.Writer, v interface{}) error {
	// Create a new JSON encoder using the provided writer.
	encoder := goccy.NewEncoder(w)

	// Encode the provided value into JSON and write it to the writer.
	return encoder.Encode(v)
}
