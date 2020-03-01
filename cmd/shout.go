package main

import (
	"errors"
	"fmt"
	"github.com/danmurf/go-cli/pkg/shout"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := &cli.App{
		Name:  "shout",
		Usage: "makes a string louder",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "level",
				Value: "1",
				Usage: "the number of !s",
			},
		},
		Action: func(c *cli.Context) error {
			text := c.Args().Get(0)
			if len(text) == 0 {
				return errors.New("nothing to shout")
			}

			level := c.Int("level")
			lm := shout.NewShout()

			fmt.Printf(lm.Shout(text, level))

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
