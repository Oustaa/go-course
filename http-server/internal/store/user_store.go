package store

import "database/sql"

type password struct {
	plainText string
	hash      string
}

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"user_name"`
	Email     string `json:"email"`
	password  `json:"-"`
	Bio       string `json:"bio"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}

type PostgresUserStore struct {
	db *sql.DB
}

func NewPostgresUserStore(db *sql.DB) *PostgresUserStore {
	return &PostgresUserStore{
		db,
	}
}

type UserStore interface {
	CreateUser(*User) error
	GetuserByUsername(username string) (*User, error)
	UpdateUser(*User) error
}

func (s *PostgresUserStore) CreateUser(user *User) error {
	query := `
		INSERT INTO users (username, bio, email password_hash)
		VALUES($1, $2, $3, $4)
		RETURNING is, create_at, updated_at
	`

	err := s.db.QueryRow(query, user.Username, user.Bio, user.Email, user.password.hash).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}
