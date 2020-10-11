package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"ostercard/internal/card"
	"ostercard/internal/command"
	"ostercard/internal/fare"
	"ostercard/internal/journey"
	"ostercard/internal/station"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		ExecuteFile(flag.Args()[0])
		return
	}

	interactiveSession()
}
func interactiveSession() {
	fmt.Println("Welcome to the oster card")
	card := card.New(30.0)
	station := station.New()
	fare := fare.New(station, card)
	journey := journey.New(fare)
	for {
		args, err := readCommands()
		if err != nil {
			fmt.Println(err.Error())
		}
		cmd := command.NewCommand(args[0], args[1:], journey, card)
		err = cmd.Execute()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func readCommands() ([]string, error) {
	reader := bufio.NewReader(os.Stdin)
	cmdString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	cmdString = strings.TrimSuffix(cmdString, "\n")
	args := strings.Split(cmdString, " ")
	return args, err
}

func ExecuteFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	card := card.New(30.0)
	station := station.New()
	fare := fare.New(station, card)
	journey := journey.New(fare)
	for scanner.Scan() {
		cmdString := scanner.Text()
		cmdString = strings.TrimSuffix(cmdString, "\n")
		args := strings.Split(cmdString, " ")

		cmd := command.NewCommand(args[0], args[1:], journey, card)
		err = cmd.Execute()
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
