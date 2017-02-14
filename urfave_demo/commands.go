package deploy

// GOLANG
//***********************************************
//
//      Filename: commands.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-02-13 15:28:42
// Last Modified: 2017-02-14 14:24:07
//***********************************************

import (
	//"encoding/json"
	"bufio"
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"io"
	//"io/ioutil"
	//"k8s.io/client-go/pkg/api/v1"
	//"k8s.io/client-go/rest"
	"os"
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
	//获取全局Flag
	fmt.Printf("global flag names :%s,%s\n", ctx.GlobalFlagNames(), ctx.GlobalString("kc"))
	//获取子命令Args
	fmt.Println("Arg is ", ctx.NArg(), ctx.Args())
	for i := 0; i < ctx.NArg(); i++ {
		fmt.Printf("Arg %d is %s\n", i, ctx.Args().Get(i))
	}
	//获取子命令Flag
	fmt.Println("Flag is ", ctx.NumFlags(), ctx.FlagNames())
	for _, flagName := range ctx.FlagNames() {
		fmt.Printf("Flag: %s, Value :%s\n", flagName, ctx.String(flagName))
	}
	fd, err := os.OpenFile(ctx.String("file"), os.O_RDONLY, os.ModeTemporary)
	defer fd.Close()
	if err != nil {
		return errors.New(fmt.Sprintf("Open file %s error !", ctx.String("file")))
	}
	//使用io/ioutil读取文件内容,据说这种速度最快
	//content, err := ioutil.ReadAll(fd)
	//if err != nil {
	//	return errors.New(fmt.Sprintf("ioutil readall from %s error !", ctx.String("file")))
	//}
	//fmt.Printf("`io/ioutil` content is :\n%s\n", string(content))
	//使用原生read和buf
	chunks := make([]byte, 1024, 1024)
	buf := make([]byte, 1024)
	//for {
	//	n, err := fd.Read(buf)
	//	if err != nil && err != io.EOF {
	//		return errors.New(fmt.Sprintf("error is %s", err))
	//	}
	//	if n == 0 {
	//		break
	//	}
	//	chunks = append(chunks, buf[:n]...)
	//}
	//fmt.Printf("`read/buf` content is :\n%s\n", string(chunks))
	//使用bufio读取
	r := bufio.NewReader(fd)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return errors.New(fmt.Sprintf("error is %s", err))
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)

	}
	fmt.Printf("`bufio` content is :\n%s\n", string(chunks))
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
