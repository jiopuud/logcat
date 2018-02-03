package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"errors"
	"encoding/json"
	"bufio"
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

func writeJson(file_path string)  {
	//var file_path = "/tmp/test_json"
	var json_obj  = make(map[string]interface{})
	var child_json = make(map[string]interface{})
	child_json["child"] = 1
	json_obj["ret"] = 0
	json_obj["test"] = child_json
	var json_str, encode_err = json.Marshal(json_obj)
	if (encode_err != nil) {
		fmt.Println(encode_err)
	}
	var file_handle, err = os.Create(file_path)
	if (err != nil) {
		fmt.Println(err)
	}
	defer file_handle.Close()
	var write = bufio.NewWriter(file_handle)
	write.Write(json_str)
	write.WriteString("\n")
	write.Flush()
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