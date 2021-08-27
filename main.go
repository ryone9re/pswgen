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
		i, err := strconv.ParseUint(c.Args()[0], 10, 64)
		if err != nil {
			log.Fatalln("Invalid arguments")
		}
		random, _ := MakeRandomStr(i, c.Bool("symbol"))
		fmt.Println(random)
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
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
