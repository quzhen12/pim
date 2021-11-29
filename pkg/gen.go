package pkg

import (
	"fmt"
	"io/ioutil"
	"path"

	"github.com/quzhen12/pim/driver"
)

func init() {
	driver.Register("gen", &gen{})
}

type gen struct{}

func (e *gen) Run(p driver.Pim) error {
	tmp := newTemplate(p)
	fmt.Println("success")
	return ioutil.WriteFile(tmp.createFilePath(), tmp.getFileTemplate(), 0700)
}

type templates struct {
	p       driver.Pim
	cmdName string
	pkgName string
	dir     string
}

func newTemplate(p driver.Pim) *templates {
	t := &templates{p: p}
	t.init()
	return t
}
func (t *templates) init() {
	t.cmdName = t.p.Params().Get(0).String()
	t.pkgName = "pkg"
	t.dir = t.p.Params().Get(1).String()
	if t.dir != "" {
		t.pkgName = path.Base(t.dir)
	}

}
func (t *templates) createFilePath() string {
	return path.Join(t.dir, fmt.Sprintf("%s.go", t.cmdName))
}

func (t *templates) getFileTemplate() []byte {
	return []byte(template(t.pkgName, t.cmdName))
}

func template(pkg, name string) string {
	return fmt.Sprintf(`package %s

import (
	"fmt"

	"github.com/quzhen12/pim/driver"
)

func init() {
	driver.Register("%s", &%s{})
}

type %s struct {}

func (%s *%s) Run(p driver.Pim) error {
	fmt.Println("Please implement the \"%s\" command.")
	return nil
}
`, pkg, name, name, name, string(name[0]), name, name)
}
