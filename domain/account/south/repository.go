package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
)

/**
 * @brief: account 仓储服务
 */
type Repository interface {
	/**
	 * @brief: 创建账户
	 * @param: entity.AccountEntity——账户充血模型
	 */
	CreateAccount(ctx context.Context, userEntity *entity.AccountEntity) error
	/**
	 * @brief: 查询账户
	 * @param: entity.AccountVo——账户 vo
	 * @return: entity.AccountEntity——账户充血模型
	 */
	GetAccount(ctx context.Context, vo *entity.AccountVo) (*entity.AccountEntity, error)

	Transaction(ctx context.Context, payer *entity.AccountEntity, payee *entity.AccountEntity, tx *entity.TransactionEntity) error
}
