package driver

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type cmder interface {
	Run(p Pim) error
}

type Pim interface {
	Params() Params
}

type mPim struct {
	name   string
	params Params
}

func newPim() *mPim {
	return &mPim{}
}

func (m *mPim) Params() Params {
	return m.params
}

func (m *mPim) parse() {
	p := os.Args
	if len(p) < 2 {
		return
	}
	m.name = p[1]
	var mm mParams
	mm.data = os.Args[2:]
	m.params = mm
}

type Params interface {
	Get(i int) SimpleParams
	Len() int
	ForEach(func(int, SimpleParams))
}

type mParams struct {
	data []string
}

func (m mParams) Get(i int) SimpleParams {
	if i >= len(m.data) {
		return mString("")
	}
	return mString(m.data[i])
}

func (m mParams) Len() int {
	return len(m.data)
}

func (m mParams) ForEach(fn func(int, SimpleParams)) {
	for i, v := range m.data {
		fn(i, mString(v))
	}
}

type mString string
type SimpleParams interface {
	Int() int
	String() string
}

func (m mString) Int() int {
	i, _ := strconv.Atoi(string(m))
	return i
}

func (m mString) String() string {
	return string(m)
}

var commands = map[string]cmder{}

func Register(name string, c cmder) {
	if _, ok := commands[name]; !ok {
		commands[name] = c
	}
}

func getCommand(cmdName string) (cmder, error) {
	cmd, ok := commands[cmdName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("pim: Command '%s' not found", cmdName))
	}
	return cmd, nil
}

type defaultCmd struct{}

func NewCommand() *defaultCmd {
	return &defaultCmd{}
}

func (d *defaultCmd) Run() error {
	p := newPim()
	p.parse()
	c, err := getCommand(p.name)
	if err != nil {
		return err
	}
	return c.Run(p)
}
