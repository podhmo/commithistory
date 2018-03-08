package commithistory

import (
	"io"

	"github.com/podhmo/commithistory/file"
)

// Parsable :
type Parsable interface {
	Parse(record []string) error
	Match(record []string, alias string) bool
}

// Unparsable :
type Unparsable interface {
	Unparse(w io.Writer) error
}

// LoadFile :
func LoadFile(filename string, ob Parsable, alias string) error {
	f, err := file.LoadFile(filename, ob.Parse, ob.Match)
	if err != nil {
		return err
	}
	defer f.Close()
	return f.Find(alias)
}

// SaveFile :
func SaveFile(filename string, ob Unparsable) error {
	return file.SaveFile(filename, ob.Unparse)
}
