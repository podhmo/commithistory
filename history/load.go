package history

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
	Match func([]string, string) bool
	Close func() error
}

// loadFile :
func loadFile(filename string, parse func([]string) error, match func([]string, string) bool) (*Finder, error) {
	fp, err := os.Open(filename)
	if err != nil {
		return nil, errors.Wrap(err, "open")
	}
	f := &Finder{r: csv.NewReader(fp), Close: fp.Close, Parse: parse, Match: match}
	return f, nil
}

// Find :
func (f *Finder) Find(alias string) error {
	for {
		record, err := f.r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return errors.Wrap(err, "csv readline")
		}
		if !f.Match(record, alias) {
			continue
		}
		return f.Parse(record)
	}
	return io.EOF
}
