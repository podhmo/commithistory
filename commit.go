package commithistory

import (
	"encoding/csv"
	"io"
	"time"

	"github.com/pkg/errors"
)

// Commit :
type Commit struct {
	ID        string
	CreatedAt time.Time
	Alias     string
	Action    string
}

func (c *Commit) match(alias string, i int) func(record []string) bool {
	return MatchByIndex(alias, i)
}

// Parse :
func (c *Commit) Parse(xs []string) error {
	if len(xs) < 4 {
		return errors.Errorf("too few %q", xs)
	}
	c.ID = xs[0]
	c.Alias = xs[1]
	ctime, err := time.Parse(time.RFC3339, xs[2])
	if err != nil {
		return err
	}
	c.CreatedAt = ctime
	c.Action = xs[3]
	return nil
}

// Unparse :
func (c *Commit) Unparse(w io.Writer) error {
	csvwriter := csv.NewWriter(w)
	row := []string{
		c.ID,
		c.Alias,
		c.CreatedAt.Format(time.RFC3339),
		c.Action,
	}
	if err := csvwriter.Write(row); err != nil {
		return err
	}
	csvwriter.Flush()
	return nil
}
