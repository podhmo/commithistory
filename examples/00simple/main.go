package main

import (
	"log"
	"time"

	"github.com/k0kubun/pp"
	"github.com/podhmo/commithistory"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ob := commithistory.Commit{
		ID:        "10f127006f06f9448eb0b86ae85aeca4",
		Alias:     "head",
		CreatedAt: time.Now(),
		Action:    "create",
	}
	filename := "./history.csv"
	if err := commithistory.SaveFile(filename, &ob); err != nil {
		return err
	}
	var ob2 commithistory.Commit
	if err := commithistory.LoadFile(filename, &ob2, commithistory.MatchByIndex("head", 1)); err != nil {
		return err
	}
	pp.Println(ob2)
	return nil
}
