package main

import (
	"fmt"
	"quzhen12/pim/driver"
	_ "quzhen12/pim/pkg"
)

func main() {
	cmd := driver.NewCommand()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
