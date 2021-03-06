package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/owtf/health_monitor/utils"

	"github.com/fatih/color"
)

var (
	cliFunctions map[string]func([]string) error
)

func init() {
	cliFunctions = make(map[string]func([]string) error)
	cliFunctions["exit"] = exit
	cliFunctions["disable"] = disableModule
	cliFunctions["enable"] = enableModule
	cliFunctions["help"] = help
	cliFunctions["status"] = status
	cliFunctions["load"] = loadProfile
	cliFunctions["disk"] = manageDisk
	cliFunctions["owtf"] = manageOWTF
	cliFunctions["profile"] = manageProfile
}

// Run will start the command line interface for the health_monitor
func Run() {
	reader := bufio.NewReader(os.Stdin)
	cyan := color.New(color.FgCyan)
	for {
		cyan.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			utils.Perror(err.Error())
		}
		text = strings.TrimSpace(text)
		command := strings.Split(text, " ")
		if len(command[0]) == 0 {
			continue
		}
		manageCommand(command)
	}
}

func manageCommand(command []string) {
	if function, ok := cliFunctions[command[0]]; ok {
		if err := function(command[1:]); err != nil {
			color.Red(err.Error())
		}
	} else {
		color.Red(fmt.Sprintf("'%s', command not found. Use 'help' to know more.",
			command[0]))
	}
}
