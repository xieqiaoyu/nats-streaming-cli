package cmd

import (
	"fmt"
)

type Resolver interface {
	Resolve(...string)
}

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
	fmt.Println("Sorry, I don't understand.")
}

func fakeCmd(name string) Cmd {
	return Cmd(func() {
		fmt.Printf("cmd %s is not a valid Resolver or a valid callable cmd func please check your code\n", name)
	})
}

type SubCmdResolver struct {
	cmds map[string]Resolver
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

func NewSubCmdResolver(cmdMap CmdMap) *SubCmdResolver {
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

	subCmdResolver := &SubCmdResolver{
		cmds: subcmds,
	}
	return subCmdResolver
}
