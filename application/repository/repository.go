package repository

import (
	"accountabully/application/bullier"
	"errors"
)

var (
	errUpdateFailed = errors.New("update failed")
	errDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Migrate(addStarterData bool) error
	InsertRule(rule bullier.Rule) error
	GetAllRules() (bullier.Rules, error)
	UpdateRule(rule bullier.Rule) error
	UpdateAllRules(rules bullier.Rules) error
	DeleteRule(id int) error
}
