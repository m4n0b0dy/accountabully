package repository

import (
	"accountabully/application/bullier"
)

type TestRepository struct{}

func NewTestRepository() *TestRepository {
	return &TestRepository{}
}

func (r *TestRepository) Migrate(addStarterData bool) error {
	return nil
}

func (r *TestRepository) InsertRule(rule bullier.Rule) error {
	return nil
}

func (r *TestRepository) GetAllRules() (bullier.Rules, error) {
	return nil, nil
}

func (r *TestRepository) UpdateRule(rule bullier.Rule) error {
	return nil
}

func (r *TestRepository) UpdateAllRules(rules bullier.Rules) error {
	return nil
}

func (r *TestRepository) DeleteRule(id int) error {
	return nil
}
