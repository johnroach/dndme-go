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
	var dice_input []string= strings.Split(input[0], "d")
	minimum_roll, minimum_roll_err := strconv.Atoi(dice_input[0])
	maximum_roll, maximum_roll_err := strconv.Atoi(dice_input[1])
	if minimum_roll_err != nil && maximum_roll_err != nil {
		fmt.Println("You probably didn't enter a valid number for the dice. :(")
	} else {
		fmt.Println(random(minimum_roll, maximum_roll * minimum_roll))
	}
}

func help(input []string) {
	fmt.Printf("%v", input)
}

func main() {
	// commands is a mapping of strings to functions
	commands := map[string]Command{
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

	prompt_input := strings.Split(prompt.Input("> ", completer), " ")
	base_command := prompt_input[0]
	_, ok := commands[base_command]
	if ok {
		commands[base_command].Fn(append(prompt_input[:0], prompt_input[1:]...))
	} else {
		fmt.Println("key not found")
	}
}
