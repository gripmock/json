//go:build !amd64

package json

import (
	goccy "github.com/goccy/go-json"
)

// MarshalIndent encodes the provided value into JSON format with indentation.
// It returns the JSON encoded data and an error if any error occurs during the encoding process.
//
// Parameters:
// - v: The value to be encoded into JSON format.
// - prefix: The prefix to be added at the beginning of each line of indented JSON data.
// - indent: The string used for indentation.
//
// Returns:
// - []byte: The JSON encoded data.
// - error: An error if any error occurs during the encoding process, otherwise nil.
func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	// Encode the provided value into indented JSON format.
	return goccy.MarshalIndent(v, prefix, indent)
}
