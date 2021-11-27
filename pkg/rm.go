package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"quzhen12/pim/driver"
)

func init() {
	driver.Register("rm", &rm{})
}

type rm struct{}

func (r *rm) Run(p driver.Pim) error {
	home, _ := os.UserHomeDir()
	f := path.Join(home, buildFilePath, "pkg", fmt.Sprintf("%s.go", p.Params().Get(0).String()))
	err := os.Remove(f)
	if err != nil {
		return err
	}
	r.build()
	return nil
}

func (r *rm) build() {
	home, _ := os.UserHomeDir()
	f := path.Join(home, buildFilePath, "build.sh")
	extc := exec.Command("bash", f)
	err := extc.Run()
	if err != nil {
		fmt.Println(err)
	}
}
