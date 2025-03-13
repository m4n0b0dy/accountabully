package repository

import (
	"accountabully/application/bullier"
	"testing"
)

func TestSQLiteRepository_Migrate(t *testing.T) {
	err := testRepo.Migrate(false)
	if err != nil {
		t.Errorf("Error migrating database: %v", err)
	}
}

func TestSQLiteRepository_InsertRule(t *testing.T) {
	rule := bullier.Rule{
		Name:   "Test Rule",
		Action: "Test Action",
		Active: true,
	}
	err := testRepo.InsertRule(rule)
	if err != nil {
		t.Errorf("Error inserting rule: %v", err)
	}
}

func TestSQLiteRepository_GetAllRules(t *testing.T) {
	rules, err := testRepo.GetAllRules()
	if err != nil {
		t.Errorf("Error getting all rules: %v", err)
	}
	if len(rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(rules))
	}
}

func TestSQLiteRepository_UpdateRule(t *testing.T) {
	rules, err := testRepo.GetAllRules()
	if err != nil {
		t.Errorf("Error getting all rules: %v", err)
	}
	if len(rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(rules))
	}

	rule := (rules)[0]
	rule.Name = "Updated Rule"
	rule.Action = "Updated Action"
	rule.Active = false

	err = testRepo.UpdateRule(rule)
	if err != nil {
		t.Errorf("Error updating rule: %v", err)
	}

	rules, err = testRepo.GetAllRules()
	if err != nil {
		t.Errorf("Error getting all rules: %v", err)
	}
	if len(rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(rules))
	}
	if (rules)[0].Name != "Updated Rule" {
		t.Errorf("Expected name to be 'Updated Rule', got '%s'", (rules)[0].Name)
	}
	if (rules)[0].Action != "Updated Action" {
		t.Errorf("Expected action to be 'Updated Action', got '%s'", (rules)[0].Action)
	}
	if (rules)[0].Active != false {
		t.Errorf("Expected active to be false, got %t", (rules)[0].Active)
	}
}

func TestSQLiteRepository_DeleteRule(t *testing.T) {
	rules, err := testRepo.GetAllRules()
	if err != nil {
		t.Errorf("Error getting all rules: %v", err)
	}
	if len(rules) != 1 {
		t.Errorf("Expected 1 rule, got %d", len(rules))
	}

	err = testRepo.DeleteRule((rules)[0].ID)
	if err != nil {
		t.Errorf("Error deleting rule: %v", err)
	}

	rules, err = testRepo.GetAllRules()
	if err != nil {
		t.Errorf("Error getting all rules: %v", err)
	}
	if len(rules) != 0 {
		t.Errorf("Expected 0 rules, got %d", len(rules))
	}
}
