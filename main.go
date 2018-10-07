package main

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Go AES"
	app.Usage = "Application for encoding / decoding strings"
	app.Description = "Encodes and decodes strings with the AES Alogrithm (ECB mode)"
	app.Version = "1.0"
	app.Author = "JanST"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "hex",
			Usage: "Define input text as hex",
		},
		cli.IntFlag{
			Name:  "klength",
			Usage: "Set key length used to encrypt/decrypt data",
			Value: 128,
		},
		cli.IntFlag{
			Name:  "rounds",
			Usage: "Set number of rounds for encryption",
			Value: 10,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"enc"},
			Usage:   "encrypt <key> <text>",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					return errors.New("Too few arguments")
				}

				key := []byte(c.Args()[0])
				var text []byte = []byte(c.Args()[1])
				if c.GlobalBool("hex") {
					var err error
					text, err = hex.DecodeString(c.Args()[1])
					if err != nil {
						return err
					}
				}

				in := bytes.NewReader(text)
				var out bytes.Buffer

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Error:", r)
					}
				}()

				Encrypt(in, &out, key, c.GlobalInt("klength"), c.GlobalInt("rounds"))
				fmt.Println("Ciphertext:", out.String())
				fmt.Println("Ciphertext (HEX):", hex.EncodeToString(out.Bytes()))
				return nil
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"dec"},
			Usage:   "decrypt <key> <text>",
			Action: func(c *cli.Context) error {
				if len(c.Args()) < 2 {
					return errors.New("Too few arguments")
				}

				key := []byte(c.Args()[0])
				var text []byte = []byte(c.Args()[1])
				if c.GlobalBool("hex") {
					var err error
					text, err = hex.DecodeString(c.Args()[1])
					if err != nil {
						return err
					}
				}

				in := bytes.NewReader(text)
				var out bytes.Buffer

				defer func() {
					if r := recover(); r != nil {
						fmt.Println("Error:", r)
					}
				}()

				Decrypt(in, &out, key, c.GlobalInt("klength"), c.GlobalInt("rounds"))
				fmt.Println("Plaintext:", out.String())
				fmt.Println("Plaintext (HEX):", hex.EncodeToString(out.Bytes()))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
