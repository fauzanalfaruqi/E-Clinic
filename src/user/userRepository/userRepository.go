package userRepository

import (
	"avengers-clinic/model/dto/userDto"
	"avengers-clinic/src/user"
	"database/sql"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) user.UserRepository {
	return &userRepository{db}
}

func (repository *userRepository) GetAllTrash() ([]userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE deleted_at IS NOT NULL ORDER BY created_at DESC;
	`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	actions, err := scanUsers(rows)
	return actions, err
}

func (repository *userRepository) GetAll() ([]userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE deleted_at IS NULL ORDER BY created_at DESC;
	`
	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	actions, err := scanUsers(rows)
	return actions, err
}

func (repository *userRepository) GetUserByID(userID string) (userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE id = $1 LIMIT 1;
	`
	user, err := scanUser(repository.db.QueryRow(query, userID))
	return user, err
}

func (repository *userRepository) GetTrashByID(userID string) (userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE id = $1 AND deleted_at IS NOT NULL LIMIT 1;
	`
	user, err := scanUser(repository.db.QueryRow(query, userID))
	return user, err
}

func (repository *userRepository) GetByID(userID string) (userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE id = $1 AND deleted_at IS NULL LIMIT 1;
	`
	user, err := scanUser(repository.db.QueryRow(query, userID))
	return user, err
}

func (repository *userRepository) GetByUsername(username string) (userDto.User, error) {
	query := `
		SELECT id, username, password, role, specialization, created_at, updated_at, deleted_at
		FROM users WHERE username = $1 AND deleted_at IS NULL LIMIT 1;
	`
	user, err := scanUser(repository.db.QueryRow(query, username))
	return user, err
}

func (repository *userRepository) Insert(user userDto.User) (string, error) {
	query := `
		INSERT INTO users (username, password, role, specialization, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;
	`
	err := repository.db.QueryRow(
		query,
		user.Username,
		user.Password,
		user.Role,
		user.Specialization,
		user.CreatedAt,
		user.UpdatedAt,
	).Scan(&user.ID);
	return user.ID, err
}

func (repository *userRepository) Update(user userDto.User) error {
	query := `
		UPDATE users SET username = $2, specialization = $3, updated_at = $4
		Where id = $1;
	`
	_, err := repository.db.Exec(
		query,
		user.ID,
		user.Username,
		user.Specialization,
		user.UpdatedAt,
	)
	return err
}

func (repository *userRepository) UpdatePassword(userId, hashPassword string) error {
	query := "UPDATE users SET password = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $1;"
	_, err := repository.db.Exec(query, userId, hashPassword)
	return err
}

func (repository *userRepository) Delete(userID string) error {
	query := "DELETE FROM users WHERE id = $1;"
	_, err := repository.db.Exec(query, userID)
	return err
}

func (repository *userRepository) SoftDelete(userID string) error {
	query := "UPDATE users SET updated_at = CURRENT_TIMESTAMP, deleted_at = CURRENT_TIMESTAMP WHERE id = $1;"
	_, err := repository.db.Exec(query, userID)
	return err
}

func (repository *userRepository) Restore(userID string) error {
	query := "UPDATE users SET updated_at = CURRENT_TIMESTAMP, deleted_at = NULL WHERE id = $1;"
	_, err := repository.db.Exec(query, userID)
	return err
}

func (repository *userRepository) IsUsernameExists(username string) bool {
	count, query := 0, "SELECT COUNT(*) FROM users WHERE username = $1;"
	repository.db.QueryRow(query, username).Scan(&count)
	return count > 0
}

func scanUser(row *sql.Row) (userDto.User, error) {
	var user userDto.User
	err := row.Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Specialization,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	return user, err
}

func scanUsers(rows *sql.Rows) ([]userDto.User, error) {
	var users []userDto.User
	for rows.Next() {
		var user userDto.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Specialization,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.DeletedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}