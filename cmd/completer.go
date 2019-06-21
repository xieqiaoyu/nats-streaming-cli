package cmd

import (
	prompt "github.com/c-bata/go-prompt"
)

type Completer interface {
	Complete(...string) []prompt.Suggest
}
