package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Author = "JJGo"
	app.Version = "1.0.0"
	app.Name = "Vagrant Uploader"
	app.Usage = "Vagrant Uploader!!"
	app.Commands = []cli.Command{
		{
			Name:    "upload",
			Aliases: []string{"u"},
			Usage:   "upload a box to vagrant cloud",
			Action: func(c *cli.Context) error {
				vagrantCloudBox := VagrantCloudBox{
					Name:     c.String("name"),
					Version:  c.String("box-version"),
					Username: c.String("username"),
					Provider: c.String("provider"),
					Token:    c.String("token"),
					Filename: c.String("filename"),
				}
				if err := vagrantCloudBox.Upload(); err != nil {
					log.Fatal(err)
				}
				return nil
			},
			Flags: []cli.Flag{
				cli.StringFlag{Name: "username, u", Usage: "Your name"},
				cli.StringFlag{Name: "name, n", Usage: "Box name"},
				cli.StringFlag{Name: "box-version, s", Usage: "Box version name"},
				cli.StringFlag{Name: "provider, p", Usage: "Provider name"},
				cli.StringFlag{Name: "token, t", Usage: "Your Vagrant cloud token"},
				cli.StringFlag{Name: "filename, f", Usage: "Box filename"},
			},
		},
	}
	app.Action = func(c *cli.Context) error {
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
