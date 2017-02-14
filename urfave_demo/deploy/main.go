// GOLANG
//***********************************************
//
//      Filename: main.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-02-13 15:11:05
// Last Modified: 2017-02-14 10:42:01
//***********************************************

package main

import (
	"fmt"
	"github.com/urfave/cli"
	deploy "github.com/xwisen/tools/urfave_demo"
	"os"
)

const version = "20170110"

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s\n", c.App.Version)
	}
	app := cli.NewApp()
	app.Name = "deployment"
	app.Version = version

	app.Flags = deploy.Flags
	app.Commands = deploy.Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		os.Exit(1)
	}
}
