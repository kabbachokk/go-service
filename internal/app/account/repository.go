package account

import (
	"database/sql"
	"fmt"
	"time"
)

type RepositoryInterface interface {
	CreateAccount(in *Input) (user *Output, err error)
	FindAccountById(pk int) (user *Output, err error)
	UpdateAccountById(pk int, in *Input) (user *Output, err error)
	DeleteAccountById(pk int) (err error)
}

type Repository struct {
	db *sql.DB
}

var _ RepositoryInterface = (*Repository)(nil)

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) CreateAccount(in *Input) (user *Output, err error) {
	var pk int
	query := fmt.Sprintf(`INSERT INTO ACCOUNTS(first_name, last_name, birth_date) VALUES('%s', '%s', '%s') RETURNING id`, in.FirstName, in.LastName, in.FormatDate())
	if err := r.db.QueryRow(query).Scan(&pk); err != nil {
		return nil, err
	}
	user, err = r.FindAccountById(pk)
	return
}

func (r *Repository) FindAccountById(pk int) (user *Output, err error) {
	var id, firstName, lastName string
	var birthDate time.Time
	query := fmt.Sprintf(`SELECT * FROM ACCOUNTS WHERE id = %d LIMIT 1`, pk)
	if err := r.db.QueryRow(query).Scan(&id, &firstName, &lastName, &birthDate); err != nil {
		return nil, err
	}
	user = &Output{
		Id:         id,
		FirstName:  firstName,
		LastName:   lastName,
		BirthYear:  birthDate.Year(),
		BirthMonth: int(birthDate.Month()),
		BirthDay:   birthDate.Day(),
	}
	return
}

func (r *Repository) UpdateAccountById(pk int, in *Input) (user *Output, err error) {
	query := fmt.Sprintf(`UPDATE ACCOUNT SET first_name='%s', last_name='%s', birth_date='%s' WHERE id=%d RETURNING id`, in.FirstName, in.LastName, in.FormatDate(), pk)
	if err := r.db.QueryRow(query).Err(); err != nil {
		return nil, err
	}
	user, err = r.FindAccountById(pk)
	return
}

func (r *Repository) DeleteAccountById(pk int) (err error) {
	query := fmt.Sprintf(`DELETE FROM USERS WHERE id=%d`, pk)
	err = r.db.QueryRow(query).Err()
	return
}
