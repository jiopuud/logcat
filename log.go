package main

import (
	"os"
	"path/filepath"
	"strings"
	"errors"
	"logcat/config"
	"bufio"
	"encoding/json"
	"io"
	"fmt"
)

func fileCheck(file_path string) (error) {
	var _, err = os.Stat(file_path)
	if err != nil {
		return err
	}
	var file_name = filepath.Base(file_path)
	var file_name_arr = strings.Split(file_name, ".")
	if len(file_name_arr) != 2 || file_name_arr[0] == "" {
		return errors.New("file name is error!")
	}
	return nil
}

func logFormatAndOutPut(file_path string) (int, error) {
	var file, err = os.Open(file_path)
	if (err != nil) {
		return config.BOOL_FALSE, err
	}
	defer file.Close()

	var reader = bufio.NewReader(file)
	var line = ""
	for {
		line, err = reader.ReadString('\n')
		if (err != nil) {
			if (err != io.EOF) {
				println(err)
			}
			break
		}
		line = strings.TrimRight(line, "\n")
		var json_obj []interface{}
		err = json.Unmarshal([] byte(line), &json_obj)
		if (err == nil) {
			var json_str, _ = json.MarshalIndent(json_obj, "", "\t")
			fmt.Println(string(json_str[:]))
		} else {
			//fmt.Println(err)
			fmt.Println(line)
		}

	}
	return config.BOOL_TRUE, nil
}
