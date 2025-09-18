package cli

import (
	"fmt"
	"os"
)

type Cli struct{
	args []string
}

func ParseArgs() *Cli {
	cli := &Cli{}
	cli.args = os.Args
	return cli
}

func (cli *Cli) PrintArgs() {
	fmt.Println("Parsed arguments:", cli.args[1:])
	
}