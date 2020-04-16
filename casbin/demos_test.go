package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnforce(t *testing.T) {
	ok, err := Enforce("alice", "data1", "read")
	assert.Nil(t, err)
	t.Logf("result=[%v]", ok)
}

func TestDemo1(t *testing.T) {
	enforcer, err := casbin.NewEnforcer("conf/model.conf", "conf/policy.csv") // Casbin 决策器需要有模型文件和策略文件作为参数
	assert.Nil(t, err)
	sub := "alice" // 需要访问资源的用户
	obj := "data1" // 将要被访问的资源
	act := "read"  // 该用户访问资源的具体操作
	ok, err := enforcer.Enforce(sub, obj, act)
	assert.Nil(t, err)
	if ok {
		t.Logf("%s has %s on %s", sub, act, obj)
	} else {
		t.Logf("%s dosen't has %s on %s", sub, act, obj)
	}
}
