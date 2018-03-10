package main

import (
	"log"

	"github.com/k0kubun/pp"
	"github.com/podhmo/commithistory"
)

// Config :
type Config struct {
	Message string `json:"message"`
}

func main() {
	c := commithistory.New("foo", commithistory.WithProfile("me"))
	var conf Config
	if err := c.Load("config.json", &conf); err != nil {
		log.Fatal(err)
	}

	pp.Println(conf)

	conf.Message = "hello"
	c.Save("config.json", &conf)
}
