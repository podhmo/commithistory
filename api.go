package commithistory

import (
	"io"
)

// Parsable :
type Parsable interface {
	Parse(record []string) error
}

// Unparsable :
type Unparsable interface {
	Unparse(w io.Writer) error
}

// LoadFile :
func LoadFile(filename string, ob Parsable, match func(record []string) bool) error {
	f, err := loadFile(filename, ob.Parse)
	if err != nil {
		return err
	}
	return f.Find(match)
}

// SaveFile :
func SaveFile(filename string, ob Unparsable) error {
	return saveFile(filename, ob.Unparse)
}
