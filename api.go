package commithistory

import (
	"path/filepath"

	"github.com/podhmo/commithistory/config"
	"github.com/podhmo/commithistory/history"
)

// todo: --profile

// Config :
type Config struct {
	*config.Config
	HistoryFile string
}

// New :
func New(name string, filename string) *Config {
	return &Config{Config: config.New(name), HistoryFile: filename}
}

// LoadCommit :
func (c *Config) LoadCommit(alias string, ob history.Parsable) error {
	dirpath, err := c.Dir(c.Name)
	if err != nil {
		return err
	}
	path := filepath.Join(dirpath, c.HistoryFile)
	return history.LoadFile(path, ob, alias)
}

// SaveCommit :
func (c *Config) SaveCommit(ob history.Unparsable) error {
	dirpath, err := c.Dir(c.Name)
	if err != nil {
		return err
	}
	path := filepath.Join(dirpath, c.HistoryFile)
	return history.SaveFile(path, ob)
}
