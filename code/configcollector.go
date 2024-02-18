package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
	// "github.com/scrapli/scrapligo/driver/options"
	// "github.com/scrapli/scrapligo/platform"
	"golang.org/x/term"
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

func getCreds() (string, string) {

	// Get username and password from the user
	var username string
	fmt.Println("Please enter your username: ")
	fmt.Scan(&username)

	fmt.Print("Enter Password: ")
	// Read password from terminal without echoing it back
	password, error := term.ReadPassword(0)
	
	if error != nil {
		log.Fatal(error)
	}
	// Turn password from bytes into string
	passwordStr := string(password)

	return username, passwordStr

}

func connectAndRunCmds {
	p, err := platform.NewPlatform("juniper_junos", "172.29.151.5", options.WithAuthNoStrictKey(), options.withAuthUsername("lab"), options.withAuthPassword("2022Nebul4"))

	if err != nil {
		return "",fmt.Errorf("failed to create platform; error: %+v", err)
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		return "",fmt.Errorf("failed to get network driver from platform %+v", err)
	}

	err = d.Open()
	if err != nil {
		return "",fmt.Errorf("failed to connect to device %+v", err)
	}

	defer d.Close()

	output, err = d.Channel.SendInput("show version")
	if err != nil {
		return "",fmt.Errorf("failed to send input to device %+v", err)
	}

	return string(output), nil
}

func main() {
	devices := fileToSlice("devices.txt")
	commands := fileToSlice("commands.txt")
	timeNow := getCurrentTime()
	uname, pword := getCreds()

	fmt.Println(devices)
	fmt.Println(commands)
	fmt.Println(timeNow)
	fmt.Println(uname, pword)

	output, err := connectAndRunCmds()
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Output: ", output)
	}

}