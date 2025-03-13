package repository

import (
	"accountabully/application/bullier"
	"database/sql"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{Conn: db}
}

func (r *SQLiteRepository) Migrate(addStarterData bool) error {
	_, err := r.Conn.Exec(`
		CREATE TABLE IF NOT EXISTS rules (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL DEFAULT '',
			action TEXT NOT NULL DEFAULT 'close',
			active BOOLEAN NOT NULL DEFAULT 1
		);
	`)
	if err != nil {
		return err
	}

	if !addStarterData {
		return nil
	}

	var count int
	err = r.Conn.QueryRow(`SELECT COUNT(*) FROM rules`).Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		_, err = r.Conn.Exec(`
			INSERT INTO rules (name, action, active) VALUES
			('instagram', 'close', 1),
			('facebook', 'close', 1),
			('twitter', 'warn', 1),
			('youtube', 'minimize', 1)
		`)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLiteRepository) InsertRule(rule bullier.Rule) error {
	_, err := r.Conn.Exec(`
		INSERT INTO rules (name, action, active)
		VALUES (?, ?, ?);
	`, rule.Name, rule.Action, rule.Active)
	return err
}

func (r *SQLiteRepository) GetAllRules() (bullier.Rules, error) {
	rows, err := r.Conn.Query(`
		SELECT id, name, action, active
		FROM rules;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rules bullier.Rules
	for rows.Next() {
		var rule bullier.Rule
		if err := rows.Scan(&rule.ID, &rule.Name, &rule.Action, &rule.Active); err != nil {
			return nil, err
		}
		rules = append(rules, rule)
	}
	return rules, nil
}

func (r *SQLiteRepository) UpdateRule(rule bullier.Rule) error {
	_, err := r.Conn.Exec(`
		UPDATE rules
		SET name = ?, action = ?, active = ?
		WHERE id = ?;
	`, rule.Name, rule.Action, rule.Active, rule.ID)
	return err

}

func (r *SQLiteRepository) DeleteRule(id int) error {
	_, err := r.Conn.Exec(`
		DELETE FROM rules
		WHERE id = ?;
	`, id)
	return err
}

func (r *SQLiteRepository) UpdateAllRules(rules bullier.Rules) error {
	tx, err := r.Conn.Begin()
	if err != nil {
		return err
	}
	for _, rule := range rules {
		_, err := tx.Exec(`
			UPDATE rules
			SET name = ?, action = ?, active = ?
			WHERE id = ?;
		`, rule.Name, rule.Action, rule.Active, rule.ID)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()

}
