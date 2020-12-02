package service

import (
	"context"

	"github.com/stanleyhsu/Go-000/Week02/internal/dao"
	"github.com/stanleyhsu/Go-000/Week02/internal/model"
)

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	return dao.GetUserByName(name)
}
