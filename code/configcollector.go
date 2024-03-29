package main

import (
	"fmt"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"golang.org/x/term"
	"log"
	"os"
	"strings"
	"time"
)

func WriteStringToFile(filename, data string) error {
	// Write the string data to the file named 'filename'
	err := os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return err
	}
	return nil
}
func connectAndRunCmds(device_type string, hostName string, username string, password string, commands []string) (string, error) {

	p, err := platform.NewPlatform(device_type, hostName, options.WithAuthNoStrictKey(), options.WithAuthUsername(username), options.WithAuthPassword(password))
	if err != nil {
		return "", fmt.Errorf("failed to create platform %+v", err)
	}

	d, err := p.GetNetworkDriver()
	if err != nil {
		return "", fmt.Errorf("failed to fetch network driver from the platform; error: %+v\n", err)
	}

	err = d.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open driver %+v", err)
	}

	defer d.Close()

	all_output := ""

	for _, cmd := range commands {
		output, err := d.Channel.SendInput(cmd)
		if err != nil {
			return "", fmt.Errorf("failed to send input to device %+v", err)
		}
		all_output += cmd + "\n"
		all_output += "-----------------------------------\n"
		all_output += string(output) + "\n"
		all_output += "-------------------------------------------------------------------\n"

	}

	file_err := WriteStringToFile(hostName+"_"+getCurrentTime()+".txt", all_output)
	if file_err != nil {
		return "", fmt.Errorf("failed to write to file %+v", err)
	}
	return all_output, nil
}

func fileToSlice(file string) []string {

	// Use ReadFile function to get content of file
	content, error := os.ReadFile(file)

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
	fmt.Print("Please enter your username: ")
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

func main() {
	devices := fileToSlice("devices.txt")
	commands := fileToSlice("commands.txt")
	uname, pword := getCreds()

	for _, device := range devices {

		_, err := connectAndRunCmds("juniper_junos", device, uname, pword, commands)
		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
