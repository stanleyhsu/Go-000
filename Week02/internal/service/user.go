package service

import (
	"context"

	"github.com/pkg/errors"

	"github.com/stanleyhsu/Go-000/Week02/internal/dao"
	"github.com/stanleyhsu/Go-000/Week02/internal/model"
)

func GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := dao.GetUserByName(name)
	if err!=nil{
		if errors.Is(err, code.ErrNotFound) {
			return return user, errors.WithMessage(err, "service no data")
		} eles {
			return return user, errors.WithMessage(err, "service no data")
		}
	}
	return user, nil
}
