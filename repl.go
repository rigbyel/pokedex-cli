package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type CliCommand struct {
	name string
	description string
	onCommandFunc func(*Config) error
}

func getCommands() map[string]CliCommand {
	return map[string]CliCommand {
		"help": {
			name: "help",
			description: "prints help info",
			onCommandFunc: callbackHelp,
		},
		"exit": {
			name: "exit",
			description: "exit programm",
			onCommandFunc: callbackExit,
		},
		"map": {
			name: "map",
			description: "show next 20 location areas",
			onCommandFunc: callbackMap,
		},
		"mapb": {
			name: "mapb",
			description: "show previous 20 location areas",
			onCommandFunc: callbackMapb,
		},
	}

}

func startREPL(cfg *Config) {
	commandsMap := getCommands()
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		userInput := scanner.Text()
		cleaned := cleanInput(userInput)

		if len(cleaned) == 0 {
			continue
		}
		commandName := cleaned[0]
		if command, ok := commandsMap[commandName]; ok {
			err := command.onCommandFunc(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Printf("No such command: %v \n", commandName)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
