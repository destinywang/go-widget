package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/sirupsen/logrus"
)

func Enforce(sub, obj, act string) (bool, error) {
	var (
		err      error
		enforcer *casbin.Enforcer
		ok       bool
	)
	if enforcer, err = casbin.NewEnforcer("conf/model.conf", "conf/policy.csv"); err != nil {
		logrus.WithError(err).Errorf("new enforcer fail")
		panic(err)
	}
	if ok, err = enforcer.Enforce(sub, obj, act); err != nil {
		logrus.WithError(err).Errorf("enforce fail")
		return false, err
	} else {
		return ok, nil
	}
}
