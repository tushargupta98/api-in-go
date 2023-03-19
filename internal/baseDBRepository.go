package internal

import (
	"github.com/jmoiron/sqlx"
)

type BaseDBRepository struct {
	db *sqlx.DB
}

func NewBaseDBRepository(db *sqlx.DB) *BaseDBRepository {
	return &BaseDBRepository{db}
}

func (r *BaseDBRepository) List(query string, args []interface{}, dest interface{}) error {
	err := r.db.Select(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseDBRepository) GetByID(query string, args []interface{}, dest interface{}) error {
	err := r.db.Get(dest, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseDBRepository) Create(query string, args []interface{}) (int, error) {
	var id int
	err := r.db.QueryRow(query, args...).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *BaseDBRepository) Update(query string, args []interface{}) error {
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (r *BaseDBRepository) Delete(query string, args []interface{}) error {
	_, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
