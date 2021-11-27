package pkg

import (
	"testing"
)

func Test_copyFile(t *testing.T) {
	a := &add{}
	err := a.copyFile("gen.go")
	if err != nil {
		t.Error(err)
	}
	t.Log()
}
