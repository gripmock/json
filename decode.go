//go:build !amd64

package json

import (
	"io"

	goccy "github.com/goccy/go-json"
)

// Decode decodes JSON data from the given io.Reader into the provided interface{}.
//
// Parameters:
// - r: The io.Reader to read the JSON data from.
// - v: The interface{} to decode the JSON data into.
//
// Returns:
// - error: An error if any error occurs during the decoding process, otherwise nil.
func Decode(r io.Reader, v interface{}) error {
	// Create a new JSON decoder using the provided io.Reader.
	decoder := goccy.NewDecoder(r)

	// Set the decoder to use numbers instead of floats.
	decoder.UseNumber()

	// Decode the JSON data from the io.Reader and store it into the provided interface{}.
	return decoder.Decode(v)
}
