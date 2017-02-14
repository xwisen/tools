package deploy

// GOLANG
//***********************************************
//
//      Filename: flags.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-02-13 15:26:02
// Last Modified: 2017-02-14 12:44:50
//***********************************************
import (
	"github.com/urfave/cli"
)

var Flags = []cli.Flag{
	cli.StringFlag{
		Name:   "kubeconfig,kc",
		Value:  "./kubeconfig",
		EnvVar: "KUBE_CONFIG",
	},
}
