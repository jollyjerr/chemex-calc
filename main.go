package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli/v2"
)

var app = cli.NewApp()

func main() {
	setup()
	commands()

	err := app.Run(os.Args)
	check(err)
}

func setup() {
	app.Name = "chemex-calc"
	app.Usage = "Figure out how much water and coffee to use in your chemex if you mess up."
	app.Version = "0.0.1"
	app.EnableBashCompletion = true
}

func commands() {
	app.Commands = []*cli.Command{
		{
			Name:    "coffee",
			Aliases: []string{"c"},
			Usage:   "Input coffee amount to find needed water amount.",
			Action:  findWaterFromCoffee,
		},
	}
}

func findWaterFromCoffee(c *cli.Context) error {
	amount, err := strconv.Atoi(c.Args().First())
	check(err)

	ans := (720 * amount) / 45
	fmt.Println(ans)

	return nil
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
