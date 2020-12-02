package dao

import (
	"database/sql"

	"github.com/pkg/errors"
	"github.com/stanleyhsu/Go-000/Week02/internal/model"
)

func GetUserByName(name string) (*model.User, error) {
	return nil, errors.Wrap(sql.ErrNoRows, "dao")
}
