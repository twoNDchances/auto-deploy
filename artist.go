package main

import "fmt"

const (
	red    = "\033[31m"
	blue   = "\033[34m"
	yellow = "\033[33m"
	green  = "\033[32m"
	purple = "\033[35m"
	reset  = "\033[0m"
)

func setColor(text, color string) string {
	switch color {
	case "red":
		return red + text + reset
	case "blue":
		return blue + text + reset
	case "green":
		return green + text + reset
	case "yelow":
		return yellow + text + reset
	case "purple":
		return purple + text + reset
	}
	return text
}

func printReport() {
	fmt.Println(setColor("All done!", "green"))
	fmt.Println("Reports:", setColor(fmt.Sprint(success), "green"), "/", setColor(fmt.Sprint(fail), "red"))
	for _, report := range finalReport {
		for key, value := range report {
			fmt.Printf("%s: %s\n", key, value)
		}
	}
}
