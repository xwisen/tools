package main

// GOLANG
//***********************************************
//
//      Filename: readline.go
//
//        Author: xwisen 1031649164@qq.com
//   Description: ---
//        Create: 2017-02-14 14:24:24
// Last Modified: 2017-02-14 14:45:13
//***********************************************

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLine(filePath string, hookfn func([]byte)) error {
	fd, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fd.Close()
	r := bufio.NewReader(fd)
	for {
		line, err := r.ReadBytes('\n')
		hookfn(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
	return nil
}
func processLine(line []byte) {
	os.Stdout.Write(line)
	fmt.Println("\n")
}

func main() {
	//ReadLine("./tm.yaml", processLine)
	fd, err := os.Open("./tm.yaml")
	if err != nil {
		fmt.Printf("%s", err)
	}
	defer fd.Close()
	r := bufio.NewReader(fd)
	for {
		line, isPrefix, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				fmt.Printf("finish readline\n")
				return
			}
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("%t:::::%s\n", isPrefix, string(line))
	}
}
