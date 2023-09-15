package main

import (
	"github.com/YangTaeyoung/aws-secret-gen/aws_secret_gen"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strings"
)

func AWSSecretGenAction(c *cli.Context) error {
	asg := aws_secret_gen.NewAWSSecretGen(
		strings.ReplaceAll(c.String("config-path"), "~", os.Getenv("HOME")),
		strings.ReplaceAll(c.String("output-path"), "~", os.Getenv("HOME")),
	)

	if err := asg.AWSInit(); err != nil {
		return err
	}
	if err := asg.Generate(); err != nil {
		return err
	}

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "aws-secret-gen"
	app.Usage = "Generate AWS Secrets"
	app.Version = "0.0.1"
	app.Description = "Super Easily get local config from AWS Secrets Manager"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:     "config-path",
			Aliases:  []string{"c"},
			Value:    "~/.aws-secret-gen/config",
			Required: false,
		},
		&cli.StringFlag{
			Name:     "output-path",
			Aliases:  []string{"o"},
			Required: true,
		},
	}

	app.Action = AWSSecretGenAction

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
