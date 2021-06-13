package main

import (
	"errors"
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
		Commands: []*cli.Command{
			{
				Name:  "fish",
				Usage: "convert fish history to zsh history",
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if path == "" {
						homeDir, err := os.UserHomeDir()
						if err != nil {
							return err
						}
						path = homeDir + "/.local/share/fish/fish_history"
					}

					b, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					hist, err := zhistconv.ParseFishHistory(b)
					if err != nil {
						return err
					}
					fmt.Print(string(hist))
					return nil
				},
			},
			{
				Name:  "parse",
				Usage: "make zsh_history readable",
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if path == "" {
						return errors.New("please specify zsh history file path")
					}
					b, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					hist := zhistconv.ParseZshHistory(b)
					fmt.Print(string(hist))
					return nil
				},
			},
			{
				Name:  "reverse",
				Usage: "convert parsed zsh_history to original",
				Action: func(c *cli.Context) error {
					path := c.Args().Get(0)
					if path == "" {
						return errors.New("please specify zsh history file path")
					}
					b, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					hist := zhistconv.ConvertToZshHistory(b)
					fmt.Print(string(hist))
					return nil
				},
			},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "write",
				Usage:   "overwrite zsh history file (default: sysout)",
				Aliases: []string{"w"},
				Value:   false,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
