package main

import (
	"fmt"
	"github.com/c-bata/go-prompt"
	"strings"
	"time"
	"math/rand"
	"strconv"
)

type Command struct {
	Fn          func([]string)
	CommandText string
	Description string
	HelpText    string
}

// commands is a mapping of strings to functions
var commands = map[string]Command{}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "roll", Description: "Rolls dice"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

// roll will roll a dice
func roll(input []string) {
	var diceInput = strings.Split(input[0], "d")
	minimumRoll, maximumRollErr := strconv.Atoi(diceInput[0])
	maximumRoll, minimumRollErr := strconv.Atoi(diceInput[1])
	if maximumRollErr != nil && minimumRollErr != nil {
		fmt.Println("You probably didn't enter a valid number for the dice. :(")
	} else {
		fmt.Println(random(minimumRoll, maximumRoll * minimumRoll))
	}
}

func help(input []string) {
	fmt.Println(commands[input[0]].HelpText)
}

func main() {

	commands = map[string]Command{
		"roll": {
		roll,
		"roll",
		"Rolls dice",
		"roll <number of times>d<dice type>"},
		"help": {
		help,
		"help",
		"Prints out help text",
		"help <command>"},
	}

	fmt.Println("Welcome to dndme!")

	for{
		promptInput := strings.Split(prompt.Input("> ", completer), " ")
		baseCommand := promptInput[0]
		_, ok := commands[baseCommand]
		if ok {
			commands[baseCommand].Fn(append(promptInput[:0], promptInput[1:]...))
		} else {
			fmt.Println("key not found")
		}
	}
}
