package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/entity"
)

/**
 * @brief: user 仓储服务
 */
type Repository interface {
	/**
	 * @brief: 创建用户
	 * @param: entity.UserEntity——用户充血模型
	 */
	CreateUser(ctx context.Context, userEntity *entity.UserEntity) error
	/**
	 * @brief: 获取用户信息
	 * @param: userId——用户唯一 id
	 * @return: entity.UserEntity——用户充血模型
	 */
	GetUser(ctx context.Context, userId string) (*entity.UserEntity, error)
}
