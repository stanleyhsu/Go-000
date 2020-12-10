package dao

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/stanleyhsu/Go-000/Week02/internal/model"
)

func GetUserByName(name string) (*model.User, error) {
	err := getUserByName(name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("sql:%s error:[%w]", sql, code.ErrNotFound)
		}
		return nil, errors.Wrapf(code.ErrDBServer, fmt.Sprintf("query: %s error(%v)", sql, err))
	}
}
