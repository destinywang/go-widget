package casbin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnforce(t *testing.T) {
	ok, err := Enforce("alice", "data1", "read")
	assert.Nil(t, err)
	t.Logf("result=[%v]", ok)
}
