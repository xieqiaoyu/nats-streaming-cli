package cmd

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
)

//Resolver resolver can decide  next step of a command
type Resolver interface {
	Resolve(...string)
}

//Suggest suggest map
type Suggest map[string]string

//CommandMap command map
type CommandMap map[string]interface{}

//Cmd a cmd action
type Cmd func()

//Resolve implement Resolver
func (c Cmd) Resolve(...string) {
	c()
}

//ParamCmd cmd accept params
type ParamCmd func(...string)

//Resolve implement Resolver
func (c ParamCmd) Resolve(t ...string) {
	c(t...)
}

func badCmd() {
	fmt.Println("Unknown command")
}

func fakeCmd(name string) Cmd {
	return Cmd(func() {
		fmt.Printf("cmd %s is not a valid Resolver or a valid callable cmd func please check your code\n", name)
	})
}

//SubCmdResolver cmd has sub cmd
type SubCmdResolver struct {
	cmds        map[string]Resolver
	suggestions []prompt.Suggest
}

//Resolve implement Resolver
func (r *SubCmdResolver) Resolve(token ...string) {
	if len(token) < 1 {
		return
	}
	child, found := r.cmds[token[0]]
	if !found {
		badCmd()
	} else if resolver, ok := child.(Resolver); ok {
		resolver.Resolve(token[1:]...)
	} else {
		fmt.Printf("child cmd %s is not a Resolver Can not excute\n", token[0])
	}
	return
}

//Complete implement completer
func (r *SubCmdResolver) Complete(t ...string) []prompt.Suggest {
	tlen := len(t)
	if tlen == 1 {
		if t[0] == "" {
			return []prompt.Suggest{}
		}
		return prompt.FilterHasPrefix(r.suggestions, t[0], true)
	} else if tlen > 1 {
		cmd, found := r.cmds[t[0]]
		if found {
			if completer, ok := cmd.(Completer); ok {
				return completer.Complete(t[1:]...)
			}
		}
	}
	return []prompt.Suggest{}
}

//NewSubCmdResolver create a sub cmd with cmd map and suggest map
func NewSubCmdResolver(cmdMap CommandMap, suggests Suggest) *SubCmdResolver {
	subcmds := map[string]Resolver{}
	for name, v := range cmdMap {
		var resolver Resolver
		if r, ok := v.(Resolver); ok {
			resolver = r
		} else if cmd, ok := v.(func(...string)); ok {
			resolver = ParamCmd(cmd)
		} else if cmd, ok := v.(func()); ok {
			resolver = Cmd(cmd)
		} else {
			resolver = fakeCmd(name)
		}
		subcmds[name] = resolver
	}

	suggestions := []prompt.Suggest{}
	for name, descr := range suggests {
		suggestions = append(suggestions, prompt.Suggest{
			Text:        name,
			Description: descr,
		})
	}
	subCmdResolver := &SubCmdResolver{
		cmds:        subcmds,
		suggestions: suggestions,
	}
	return subCmdResolver
}
