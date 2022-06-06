package cli

import (
	"flag"
	"fmt"
)

const daemonFlag = "daemon"

var daemon = flag.Bool(daemonFlag, false, "Run daemon")

func NewPermissionCi() *PermissionsCli {
	return &PermissionsCli{
		Description: "Simple cli tool to manage permissions",
	}
}

type PermissionsCli struct {
	Description string
}

func (cli *PermissionsCli) Run() error {
	flag.Parse()

	if *help {
		info := cli.Description + "\n\n"
		flag.VisitAll(func(f *flag.Flag) {
			info += fmt.Sprintf("%11s\t\t%s\n", f.Name, f.Usage)
		})

		fmt.Println(info)
		return nil
	}

	if *daemon {
		return NewPermissionsDaemon().Listen()
	}

	return NewPermissionClient().Handle()
}
