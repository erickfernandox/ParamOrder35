package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func formatParameters(lst []string) []string {
	params := make([]string, 0)
	line := ""
	for i, item := range lst {
		param := fmt.Sprintf("%s=FUZZ", item)
		if i > 0 && i%35 == 0 {
			params = append(params, line)
			line = ""
		}
		if line != "" {
			line += "&"
		}
		line += param
	}
	if line != "" {
		params = append(params, line)
	}
	return params
}

func main() {
	listFile := flag.String("l", "", "File containing the list")
	flag.Parse()

	if *listFile == "" {
		fmt.Println("Please provide a list file using the -listfile flag.")
		os.Exit(1)
	}

	file, err := os.Open(*listFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lst []string
	for scanner.Scan() {
		lst = append(lst, scanner.Text())
	}

	params := formatParameters(lst)
	for _, param := range params {
		fmt.Printf("?%s\n", param)
	}
}
