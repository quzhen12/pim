package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/quzhen12/pim/driver"
)

func init() {
	driver.Register("add", &add{})
}

type add struct{}

const buildFilePath = "./pim/build/pim"

func (a *add) Run(p driver.Pim) error {
	a.copyFile(p.Params().Get(0).String())
	a.build()
	return nil
}

func (a *add) copyFile(filePath string) error {
	r, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}
	fileName := path.Base(filePath)
	home, _ := os.UserHomeDir()
	return ioutil.WriteFile(path.Join(home, buildFilePath, "pkg", fileName), r, 0700)
}

func (a *add) build() {
	home, _ := os.UserHomeDir()
	extc := exec.Command("bash", path.Join(home, buildFilePath, "build.sh"))
	err := extc.Run()
	if err != nil {
		fmt.Println(err)
	}
}
