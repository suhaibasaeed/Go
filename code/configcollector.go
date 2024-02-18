package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func FileToSlice(file string) []string {
	
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


func main() {
	devices := FileToSlice("devices.txt")
	commands := FileToSlice("commands.txt")
	fmt.Println(devices)
	fmt.Println(commands)
}