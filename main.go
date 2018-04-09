package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"errors"
)

func catAction(file_name string) error {
	var err = fileCheck(file_name)
	if (err != nil) {
		return err
	}
	logFormatAndOutPut(file_name)
	return nil
}

func argCheck(c *cli.Context) error {
	var arg_num = c.NArg()
	if arg_num <= 0 {
		var msg = fmt.Sprintf("arg num error, %s", c.App.UsageText)
		return errors.New(msg)
	}
	return nil
}

func main() {
	var app = cli.NewApp()
	app.Name = "logcat"
	app.Usage = "a command format log output to standrad output"
	app.Version = "1.0.0"
	app.UsageText = fmt.Sprintf("please %s log_file", app.Name)
	app.ArgsUsage = "the path that you want to format output"

	app.Action = func(c *cli.Context) error {
		var err = argCheck(c)
		if err != nil {
			fmt.Println(err)
			return err
		}
		var file_name = c.Args().Get(0)
		//writeJson(file_name)
		err = catAction(file_name)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	}

	app.Run(os.Args)
}