package tools

import (
	"testing"
)

func TestGenId(t *testing.T) {
	id, err := GenId()
	if nil != err {
		t.Error(err)
	}
	t.Log("id = ", id)
}
