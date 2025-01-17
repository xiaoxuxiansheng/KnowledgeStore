package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/entity"
)

/**
 * @brief: user domain 服务
 */
type IService interface {
	/**
	 * @brief: 注册用户
	 * @return: entity.UserEntity——用户充血模型
	 */
	Register(ctx context.Context) (*entity.UserEntity, error)
	/**
	 * @brief: 查询用户
	 * @param: userId——用户唯一 id
	 * @return: entity.UserEntity——用户充血模型
	 */
	GetUser(ctx context.Context, userId string) (*entity.UserEntity, error)
}
