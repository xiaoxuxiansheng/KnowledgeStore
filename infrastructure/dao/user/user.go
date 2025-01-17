package user

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/south"
)

type Dao struct {
}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateUser(ctx context.Context, userEntity *entity.UserEntity) error {
	return nil
}

func (d *Dao) GetUser(ctx context.Context, userId string) (*entity.UserEntity, error) {
	return nil, nil
}
