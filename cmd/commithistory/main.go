package main

import (
	"fmt"
	"go/build"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

type opt struct {
	Pkg string
}

func guessPkg() (string, error) {
	curdir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(curdir)
	if err != nil {
		return "", err
	}
	for _, srcdir := range build.Default.SrcDirs() {
		if strings.HasPrefix(path, srcdir) {
			pkgname := strings.TrimLeft(strings.Replace(path, srcdir, "", 1), "/")
			return pkgname, nil
		}
	}

	return "", errors.Errorf("%q is not subdir of srcdirs(%q)", path, build.Default.SrcDirs())
}

func main() {
	var opt opt
	app := kingpin.New("strangejson", "strangejson")
	app.Flag("pkg", "package").StringVar(&opt.Pkg)

	if _, err := app.Parse(os.Args[1:]); err != nil {
		app.FatalUsage(fmt.Sprintf("%v", err))
	}

	if opt.Pkg == "" {
		pkg, err := guessPkg()
		if err != nil {
			app.FatalUsage(fmt.Sprintf("%v", err))
		}
		opt.Pkg = pkg
		log.Printf("guess pkg name .. %q\n", opt.Pkg)
	}

	if err := run(opt); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run(opt opt) error {
	fmt.Println(opt.Pkg)
	return nil
}
