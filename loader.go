package main

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

func loadFileLines(path string) (*[]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &lines, nil
}

func loadCreds(credsFile string) (map[string]interface{}, error) {
	lines, err := loadFileLines(credsFile)
	if err != nil {
		return nil, err
	}
	creds := make(map[string]interface{})
	for _, line := range *lines {
		if strings.HasPrefix(line, "USERNAME") {
			creds["username"] = strings.ReplaceAll(line, "USERNAME=", "")
		}
		if strings.HasPrefix(line, "HOSTNAME") {
			creds["hostname"] = strings.ReplaceAll(line, "HOSTNAME=", "")
		}
		if strings.HasPrefix(line, "PORT") {
			value, err := strconv.Atoi(strings.ReplaceAll(line, "PORT=", ""))
			if err != nil {
				return nil, err
			}
			if value <= 0 {
				return nil, errors.New("port can not be negative")
			}
			creds["port"] = value
		}
		if strings.HasPrefix(line, "PASSWORD") {
			creds["password"] = strings.ReplaceAll(line, "PASSWORD=", "")
		}
		if strings.HasPrefix(line, "SITE_FILE") {
			creds["siteFile"] = strings.ReplaceAll(line, "SITE_FILE=", "")
		}
		if strings.HasPrefix(line, "COMMANDS") {
			cmds, err := loadFileLines(strings.ReplaceAll(line, "COMMANDS=", ""))
			if err != nil {
				return nil, err
			}
			creds["commands"] = *cmds
		}
	}
	return creds, nil
}
