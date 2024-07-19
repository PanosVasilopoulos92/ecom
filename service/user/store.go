package user

import (
	"database/sql"
	"errors"

	"github.com/PanosVasilopoulos92/ecom/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE email = ?", email)

	user, err := scanUniqueUserRow(row)
	if err != nil {
		return nil, errors.New("no such a user in database")
	}

	if user.ID == 0 {
		return nil, errors.New("no such a user in database")
	}

	return user, nil

}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE ID = ?", id)

	user, err := scanUniqueUserRow(row)

	if err != nil {
		return nil, errors.New("no such a user in database")
	}

	if user.ID == 0 {
		return nil, errors.New("no such a user in database")
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	return nil
}

func scanUniqueUserRow(row *sql.Row) (*types.User, error) {
	user := new(types.User)

	err := row.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}
