package main

import (
	"fmt"
	"log"
	"os"

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
				Name:    "fish",
				Usage:   "specify fish_history file",
				Aliases: []string{"f"},
			},
		},
		Action: func(c *cli.Context) error {
			fishFilePath := c.String("fish")
			hist, err := parseFishHistory(fishFilePath)
			if err != nil {
				return err
			}

			fmt.Println(hist)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

// 131は0x83の10進数表現 zsh_historyの特殊仕様
const metachar = 131
const start = 130
const end = 158
const xA0 = 160

func parseNonAscii(latin1Byte []byte) string {
	isMarking := false
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		if codePoint == metachar {
			isMarking = true
			continue
		}

		if isMarking {
			// 6bit目を反転させるために
			// 0x20をXORする
			invertCodePoint := codePoint ^ 32
			byteBuffer = append(byteBuffer, invertCodePoint)
			isMarking = false
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}

func convertToZshHistory(latin1Byte []byte) string {
	var byteBuffer []byte

	for _, codePoint := range latin1Byte {
		isInverse := false
		// 131は0metacharの10進数表現
		if (start < codePoint && codePoint < end) || codePoint == xA0 {
			isInverse = true
		}

		if isInverse {
			// 6bit目を反転させるために
			// 0x20をXORする
			invertCodePoint := codePoint ^ 32
			byteBuffer = append(byteBuffer, metachar)
			byteBuffer = append(byteBuffer, invertCodePoint)
			isInverse = false
		} else {
			byteBuffer = append(byteBuffer, codePoint)
		}
	}

	return string(byteBuffer)
}
