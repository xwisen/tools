package deploy

// GOLANG
//***********************************************
//
//      Filename: commands.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-02-13 15:28:42
// Last Modified: 2017-02-14 10:39:38
//***********************************************

import (
	"fmt"
	"github.com/urfave/cli"
)

var Commands = []cli.Command{
	{
		Name:    "create",
		Aliases: []string{"c"},
		Usage:   "create a deployment",
		Action:  CreateDeployment,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "f,fl,file",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
			cli.StringFlag{
				Name:  "a",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
			cli.StringFlag{
				Name:  "b",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
			cli.StringFlag{
				Name:  "c",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
		},
	},
	{
		Name:    "delete",
		Aliases: []string{"d"},
		Usage:   "delete a deployment",
		Action:  DeleteDeployment,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "f,file",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
		},
	},
	{
		Name:    "update",
		Aliases: []string{"u"},
		Usage:   "update a deployment",
		Action:  UpdateDeployment,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "f,file",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
		},
	},
	{
		Name:    "patch",
		Aliases: []string{"p"},
		Usage:   "patch a deployment",
		Action:  PatchDeployment,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "f,file",
				Value: "./deploy.json",
				Usage: "deployment file",
			},
		},
	},
}

func CreateDeployment(ctx *cli.Context) error {
	fmt.Println("create a deployment")
	fmt.Println("Arg is ", ctx.NArg(), ctx.Args())
	for i := 0; i < ctx.NArg(); i++ {
		fmt.Printf("Arg %d is %s\n", i, ctx.Args().Get(i))
	}

	fmt.Println("Flag is ", ctx.NumFlags(), ctx.FlagNames())
	for _, flagName := range ctx.FlagNames() {
		fmt.Printf("Flag: %s, Value :%s\n", flagName, ctx.String(flagName))
	}
	return nil
}

func DeleteDeployment(ctx *cli.Context) error {
	fmt.Println("delete a deployment")
	return nil
}

func UpdateDeployment(ctx *cli.Context) error {
	fmt.Println("update a deployment")
	return nil
}

func PatchDeployment(ctx *cli.Context) error {
	fmt.Println("patch a deployment")
	return nil
}
