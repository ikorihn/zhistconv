package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/r57ty7/zhistconv"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "f2zhistory",
		Usage: "convert fish hist to zsh hist",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "write",
				Usage:   "overwrite zsh history file (default: sysout)",
				Aliases: []string{"w"},
				Value:   false,
			},
			&cli.StringFlag{
				Name:  "fish",
				Usage: "specify fish_history file",
			},
			&cli.StringFlag{
				Name:  "parse",
				Usage: "convert zsh_history readable",
			},
			&cli.StringFlag{
				Name:  "reverse",
				Usage: "convert parsed zsh_history to original",
			},
		},
		Action: func(c *cli.Context) error {
			fishFilePath := c.String("fish")
			if fishFilePath != "" {
				fishHist, err := zhistconv.ParseFishHistory(fishFilePath)
				if err != nil {
					return err
				}
				fmt.Println(fishHist)
			}

			parse := c.String("parse")
			if parse != "" {
				b, err := ioutil.ReadFile(parse)
				if err != nil {
					return err
				}
				hist := zhistconv.ParseZshHistory(b)
				fmt.Println(hist)
			}

			reverse := c.String("reverse")
			if reverse != "" {
				b, err := ioutil.ReadFile(reverse)
				if err != nil {
					return err
				}
				hist := zhistconv.ConvertToZshHistory(b)
				fmt.Println(hist)
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
