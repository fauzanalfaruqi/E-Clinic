package actionRepository

import (
	"avengers-clinic/model/dto/actionDto"
	"avengers-clinic/src/action"
	"database/sql"
)

type actionRepository struct {
	db *sql.DB
}

func NewActionRepository(db *sql.DB) (action.ActionRepository) {
	return &actionRepository{db}
}

func (repository *actionRepository) GetAll() ([]actionDto.Action, error) {
	query := `
		SELECT id, name, price, description, created_at, updated_at
		FROM actions WHERE deleted_at IS NULL;
	`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	actions, err := scanActions(rows)
	return actions, err
}

func (repository *actionRepository) GetByID(actionID string) (actionDto.Action, error) {
	query := `
		SELECT id, name, price, description, created_at, updated_at
		FROM actions WHERE id = $1 AND deleted_at IS NULL LIMIT 1;
	`
	action, err := scanAction(repository.db.QueryRow(query, actionID))
	return action, err
}

func (repository *actionRepository) Insert(action actionDto.Action) (string, error) {
	query := `
		INSERT INTO actions (name, price, description, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;
	`
	err := repository.db.QueryRow(
		query,
		action.Name,
		action.Price,
		action.Description,
		action.CreatedAt,
		action.UpdatedAt,
	).Scan(&action.ID)
	return action.ID, err
}

func (repository *actionRepository) Update(action actionDto.Action) error {
	query := `
		UPDATE actions SET name = $2, price = $3, description = $4, updated_at = $5
		WHERE id = $1;
	`
	_, err := repository.db.Exec(
		query,
		action.ID,
		action.Name,
		action.Price,
		action.Description,
		action.UpdatedAt,
	)
	return err
}

func (repository *actionRepository) Delete(actionID string) error {
	query := "DELETE FROM actions WHERE id = $1;"
	_, err := repository.db.Exec(query, actionID)
	return err
}

func (repository *actionRepository) SoftDelete(actionID string) error {
	query := "UPDATE actions SET deleted_at = CURRENT_TIMESTAMP WHERE id = $1;"
	_, err := repository.db.Exec(query, actionID)
	return err
}

func (repository *actionRepository) Restore(actionID string) error {
	query := "UPDATE actions SET deleted_at = NULL WHERE id = $1;"
	_, err := repository.db.Exec(query, actionID)
	return err
}

func (repository *actionRepository) IsNameExist(name string) bool {
	count, query := 0, "SELECT COUNT(*) FROM actions WHERE name = $1;"
	repository.db.QueryRow(query, name).Scan(&count)
	return count > 0
}

func scanAction(row *sql.Row) (actionDto.Action, error) {
	var action actionDto.Action
	err := row.Scan(
		&action.ID,
		&action.Name,
		&action.Price,
		&action.Description,
		&action.CreatedAt,
		&action.UpdatedAt,
	)
	return action, err
}

func scanActions(rows *sql.Rows) ([]actionDto.Action, error) {
	var actions []actionDto.Action
	for rows.Next() {
		var action actionDto.Action
		err := rows.Scan(
			&action.ID,
			&action.Name,
			&action.Price,
			&action.Description,
			&action.CreatedAt,
			&action.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		actions = append(actions, action)
	}
	return actions, nil
}