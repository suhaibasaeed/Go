package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func fileToSlice(file string) []string {
	
	// Use ReadFile function to get content of file
	content, error := ioutil.ReadFile(file)

	// Check for errors, if yes then print error and exit
	if error != nil {
		log.Fatal(error)
	}

	// Convert file content into string and print
	fileStr := string(content)
	// Turn string into slice
	original_slice := strings.Split(fileStr, "\n")

	// Loop through each string in the list and remove whitespace and empty strings
	file_list := []string{}
	for _, str := range original_slice {
		// Only add non-empty strings to new slice
		if str != "" {
			trimmedStr := strings.TrimSpace(str)
			file_list = append(file_list, trimmedStr)
		}
		
	} 
	
	return file_list
}

func getCurrentTime() string {
	
	// Get current time
	currentTime := time.Now()

	formattedTime := currentTime.Format("02-01-06@15.04")

	return formattedTime
}


func main() {
	devices := fileToSlice("devices.txt")
	commands := fileToSlice("commands.txt")
	timeNow := getCurrentTime()
	fmt.Println(devices)
	fmt.Println(commands)
	fmt.Println(timeNow)
}