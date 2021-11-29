package main

import (
	"fmt"

	"github.com/quzhen12/pim/driver"
	_ "github.com/quzhen12/pim/pkg"
)

func main() {
	cmd := driver.NewCommand()
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
