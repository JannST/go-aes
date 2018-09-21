package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Version = "1.0"
	app.Author = "JanST"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "file",
			Usage: "programm will resolve input to file path and process file",
		},
		cli.IntFlag{
			Name:  "klength",
			Usage: "Sets the key length used to encrypt/decrypt data",
			Value: 34,
		},
		cli.StringFlag{
			Name:  "mode",
			Usage: "Sets encryption mode (ECB or CBC)",
			Value: "ECB",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"enc"},
			Usage:   "encrypts data",
			Action: func(c *cli.Context) error {
				fmt.Println(c.GlobalBool("file"))
				fmt.Println(c.Args())
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"dec"},
			Usage:   "decrypts data",
			Action: func(c *cli.Context) error {
				fmt.Println(c.GlobalBool("file"))
				fmt.Println(c.Args())
				fmt.Println("added task: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
