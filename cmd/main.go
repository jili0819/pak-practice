package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const (
	version = "v1.0.0"
)

var (
	configFlag = cli.StringFlag{
		Name:     "config",
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: false,
	}
	portFlag = cli.StringFlag{
		Name:     "port",
		Aliases:  []string{"p"},
		Usage:    "port 8080",
		Required: true,
	}
	stringFlag = cli.StringFlag{
		Name:     "string",
		Aliases:  []string{},
		Usage:    "input string",
		Value:    "default string",
		Required: true,
	}
	intFlag = cli.StringFlag{
		Name:     "int",
		Aliases:  []string{},
		Usage:    "input int",
		Value:    "default int",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "test cli"
	app.Version = version
	flags := []cli.Flag{
		&configFlag,
		&portFlag,
	}
	// 可以分多条命令，分别启动不同的服务

	app.Commands = []*cli.Command{
		{
			//go run main version
			Name:    "version",
			Aliases: []string{},
			Usage:   "Application version and build",
			Action:  versionCmd,
		},
		{
			//./main check
			Name:    "check",
			Aliases: []string{},
			Usage:   "Run the demo",
			Action:  check,
			Flags:   []cli.Flag{&stringFlag, &intFlag},
		},
		{
			//go run server.go run
			Name:    "run",
			Aliases: []string{},
			Usage:   "Run the demo",
			Action:  start,
			Flags:   append(flags),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func start(cliCtx *cli.Context) error {
	fmt.Println(fmt.Sprintf("cli port :%d", cliCtx.Int(portFlag.Name)))
	return nil
}

func versionCmd(cliCtx *cli.Context) error {
	fmt.Println(version)
	return nil
}

func check(cliCtx *cli.Context) error {
	fmt.Println("----check---")
	fmt.Println(fmt.Sprintf("--check---:string flag : %s", cliCtx.String(stringFlag.Name)))
	fmt.Println(fmt.Sprintf("--check---:int flag : %d", cliCtx.Int(intFlag.Name)))
	return nil
}
