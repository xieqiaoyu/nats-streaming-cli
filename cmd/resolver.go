package cmd

import (
	"fmt"
	prompt "github.com/c-bata/go-prompt"
)

type Resolver interface {
	Resolve(...string)
}

type Suggest map[string]string

type CmdMap map[string]interface{}

type Cmd func()

func (c Cmd) Resolve(...string) {
	c()
}

type CmdWithParam func(...string)

func (c CmdWithParam) Resolve(t ...string) {
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

type SubCmdResolver struct {
	cmds        map[string]Resolver
	suggestions []prompt.Suggest
}

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

func NewSubCmdResolver(cmdMap CmdMap, suggests Suggest) *SubCmdResolver {
	subcmds := map[string]Resolver{}
	for name, v := range cmdMap {
		var resolver Resolver
		if r, ok := v.(Resolver); ok {
			resolver = r
		} else if cmd, ok := v.(func(...string)); ok {
			resolver = CmdWithParam(cmd)
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
