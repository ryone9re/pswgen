package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "pswgen"
	app.Usage = "Password generator"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "symbol, s",
			Usage: "Including symbol",
		},
	}
	app.Action = func(c *cli.Context) error {
		if c.Args() == nil || len(c.Args()) == 0 {
			log.SetFlags(0)
			log.Fatal("\npswgen [OPTIONS] [PASSWORD LENGTH]\n\nOPTIONS:\n  -h, --help:\n	Show this help message and exit.\n  -v, --version:\n	Show version and exit.\n  -s, --symbol:\n	Add symbol to password.")
		}
		i, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil || i < 1 {
			log.SetFlags(0)
			log.Fatal("\npswgen [OPTIONS] [PASSWORD LENGTH]\n\nOPTIONS:\n  -h, --help:\n	Show this help message and exit.\n  -v, --version:\n	Show version and exit.\n  -s, --symbol:\n	Add symbol to password.")
		}
		random, _ := MakeRandomStr(i, c.Bool("symbol"))
		fmt.Println(random)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.SetFlags(0)
		log.Fatal("\npswgen [OPTIONS] [PASSWORD LENGTH]\n\nOPTIONS:\n  -h, --help:\n	Show this help message and exit.\n  -v, --version:\n	Show version and exit.\n  -s, --symbol:\n	Add symbol to password.")
	}
}

func MakeRandomStr(digit uint64, symbol bool) (string, error) {
	var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if symbol {
		letters += "-_/*+.,!#$%&()~|"
	}

	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return "", errors.New("unexpected error")
	}

	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}
