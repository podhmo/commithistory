package config

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

// Config :
type Config struct {
	Name string

	Dir          func() (string, error)
	NotFound     func(filepath string, err error) (io.ReadCloser, error)
	WriteDefault func(filepath string) error
	Load         func(name string, ob interface{}) error
	Unmarshal    func(r io.Reader, ob interface{}) error
	Save         func(name string, ob interface{}) error
	Marshal      func(w io.Writer, ob interface{}) error
}

// Default :
func Default(name string) *Config {
	return &Config{Name: name}
}

// Save :
func Save(c *Config, name string, ob interface{}) error {
	if c.Save != nil {
		return c.Save(name, ob)
	}
	d, err := Dir(c)
	if err != nil {
		return err
	}

	path := filepath.Join(d, name)
	log.Printf("save. %q\n", path)

	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return err
	}
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	return Marshal(c, fp, ob)
}

// Marshal :
func Marshal(c *Config, w io.Writer, ob interface{}) error {
	if c.Marshal != nil {
		return c.Marshal(w, ob)
	}
	encoder := json.NewEncoder(w)
	return encoder.Encode(ob)
}

// Load :
func Load(c *Config, name string, ob interface{}) error {
	if c.Load != nil {
		return c.Load(name, ob)
	}
	d, err := Dir(c)
	if err != nil {
		return err
	}

	path := filepath.Join(d, name)
	log.Printf("load. %q\n", path)

	var fp io.ReadCloser
	fp, err = os.Open(path)
	if err != nil {
		fp, err = NotFound(c, path, err)
		if err != nil {
			return err
		}
	}
	defer fp.Close()
	return Unmarshal(c, fp, ob)
}

// Unmarshal :
func Unmarshal(c *Config, r io.Reader, ob interface{}) error {
	if c.Unmarshal != nil {
		return c.Unmarshal(r, ob)
	}
	decoder := json.NewDecoder(r)
	return decoder.Decode(ob)
}

// Dir :
func Dir(c *Config) (string, error) {
	if c.Dir != nil {
		return c.Dir()
	}
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".config", c.Name), nil
}

// NotFound :
func NotFound(c *Config, path string, err error) (io.ReadCloser, error) {
	if c.NotFound != nil {
		return c.NotFound(path, err)
	}
	log.Printf("not found. %q\n", path)
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}
	if err := WriteDefault(c, path); err != nil {
		return nil, err
	}
	return os.Open(path)
}

// WriteDefault :
func WriteDefault(c *Config, path string) error {
	if c.WriteDefault != nil {
		return c.WriteDefault(path)
	}
	log.Printf("create. %q\n", path)
	fp, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fp.Close()
	_, rerr := fmt.Fprintln(fp, "{}")
	return rerr
}
