package internal

import (
	"encoding/json"
	"io"
)

// encode encodes FCV objects to JSON byte data.
// TODO remove JSON encoding and use binary encoding instead.
func encode(fcvObjects []FCVObject) ([]byte, error) {
	return json.Marshal(fcvObjects)
}

// decode deocdes JSON byte data to FCV objects.
// TODO remove JSON decoding and use binary decoding instead.
func decode(blob []byte) ([]FCVObject, error) {
	var fcvObjects []FCVObject
	err := json.Unmarshal(blob, &fcvObjects)
	if err != nil {
		return nil, err
	}
	return fcvObjects, nil
}

// WriteContent write FCV objects to file.
// uses io.ReadWriteCloser for testability.
func WriteContent(file io.ReadWriteCloser, fcvObjects []FCVObject) error {
	blob, err := encode(fcvObjects)
	if err != nil {
		return err
	}
	_, err = file.Write(blob)
	if err != nil {
		return err
	}
	return nil
}
