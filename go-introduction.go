package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const DEFAULT_TRACE_ATTEMPTS = 5
const DELAY = 5

func main() {

	intro()

	for {

		options()

		option := readOption()

		switch option {
		case 1:
			trace(DEFAULT_TRACE_ATTEMPTS)
		case 2:
			fmt.Println("Not Implemented")
		case 0:
			fmt.Println("Exiting")
			os.Exit(0)
		default:
			fmt.Println("Invalid option")
			os.Exit(-1)
		}
	}

}

func intro() {
	name := "Victor"
	version := 1.2
	fmt.Println("Hello,", name)
	fmt.Println("Program version:", version)
}

func options() {
	fmt.Println("1- Start Trace")
	fmt.Println("2- Logs")
	fmt.Println("0- Exit")
}

func readOption() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("You choose:", option)
	return option
}

func trace(times int) {

	fmt.Println("tracing...")

	/*
		applications := []string{"https://random-status-code.herokuapp.com/", "https://viacep.com.br/ws/01001000/json/",
			"https://viacep.com.br/invalid-uri"}
	*/

	applications := getWebApplicationsFromFile()

	for i := 0; i < times; i++ {
		for _, application := range applications {
			testWebApplication(application)
		}
		time.Sleep(DELAY * time.Second)
	}
}

func testWebApplication(applicationUrl string) {

	resp, err := http.Get(applicationUrl)

	if err != nil {
		fmt.Println("An error occurred while trying to make HTTP Request to:", applicationUrl, "=>", err)
	}

	statusCode := resp.StatusCode

	if statusCode == 200 {
		fmt.Println("Web Application:", applicationUrl, " successfully loaded. It's working correctly.")
		logger(applicationUrl, true)
	} else if statusCode == 404 {
		fmt.Println("Web Application:", applicationUrl, " doesn't exist [", statusCode, "]", "Check the URL [", applicationUrl, "]")
		logger(applicationUrl, false)
	} else {
		fmt.Println("Web Application:", applicationUrl, "is NOT working. HTTP Status Code:", resp.StatusCode)
		logger(applicationUrl, true)
	}
}

func getWebApplicationsFromFile() []string {
	var applications []string

	//file, err := os.Open("webApplications.txt") //By this way, Go returns only the file pointer.

	//file, err := ioutil.ReadFile("webApplications.txt") //it returns the file Bytes

	file, err := os.Open("webApplications.txt")

	if err != nil {
		fmt.Println("An error occurred while trying to open file:", err)
	}

	//fmt.Println(string(file)) //convert bytes to string  (when using ioutil) but it just print the file content and we need to read line by line

	scanner := bufio.NewReader(file)

	for {

		line, err := scanner.ReadString('\n') //Using \n as delimiter, the Go reads line and add \n

		line = strings.TrimSpace(line) //Removing the \n

		applications = append(applications, line)

		if err == io.EOF {
			//if we arrive the end of file, break the loop.
			break
		}
	}

	file.Close()

	return applications
}

func logger(applicationUrl string, wasOnline bool) {

	file, err := os.OpenFile("applications-tracing.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("An error has occurred while trying to Open or create log file:", err)
	}

	var content string

	if wasOnline {
		content = applicationUrl + " was online\n"
	} else {
		content = applicationUrl + " was offline\n"
	}

	if err != nil {
		fmt.Println("an error has occurred while creating log file:", err)
	}

	file.WriteString(content)
}
