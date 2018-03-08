package file

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

// Finder :
type Finder struct {
	r     *csv.Reader
	Parse func([]string) error
	Close func() error
}

// LoadFile :
func LoadFile(filename string, parse func([]string) error) (*Finder, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "open")
	}
	f := &Finder{r: csv.NewReader(fp), Close: fp.Close, Parse: parse}
	return f, nil
}

// MatchByIndex :
func MatchByIndex(alias string, i int) func(record []string) bool {
	return func(record []string) bool {
		return record[i] == alias
	}
}

// Find :
func (f *Finder) Find(match func([]string) bool) error {
	for {
		record, err := f.r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "csv readline")
		}
		if !match(record) {
			continue
		}
		return f.Parse(record)
	}
	return io.EOF
}
