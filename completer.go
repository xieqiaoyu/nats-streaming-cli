package main

import (
	prompt "github.com/c-bata/go-prompt"
	"strings"
)

var suggestions = []prompt.Suggest{
	{Text: "show", Description: "show server mesage"},
	{Text: "pub", Description: "publish msg to a channel"},
	{Text: "list", Description: "list message in a channel"},
	{Text: "exit", Description: "Exit cli"},
}

var cmdShowSuggestions = []prompt.Suggest{
	{Text: "channel", Description: "show one channel info"},
	{Text: "channels", Description: "show channels info"},
	{Text: "server", Description: "show server info"},
	{Text: "store", Description: "show store info"},
	{Text: "clients", Description: "show clients info"},
}

func completer(in prompt.Document) []prompt.Suggest {
	input := in.TextBeforeCursor()
	if input == "" {
		return []prompt.Suggest{}
	}
	blocks := strings.Split(input, " ")
	if len(blocks) == 1 {
		return prompt.FilterHasPrefix(suggestions, blocks[0], true)
	}

	switch blocks[0] {
	case "show":
		return prompt.FilterHasPrefix(cmdShowSuggestions, blocks[1], true)
	}
	return []prompt.Suggest{}
}
