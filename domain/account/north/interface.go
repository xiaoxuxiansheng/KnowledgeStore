package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
)

/**
 * @brief: account domain 服务
 */
type IService interface {
	/**
	 * @brief: 创建账户
	 * @param: entity.AccountEntity——账户充血模型
	 */
	CreateAccount(ctx context.Context, account *entity.AccountEntity) error
	/**
	 * @brief: 查询账户
	 * @param: entity.AccountVo——账户 vo
	 * @return: entity.AccountEntity——账户充血模型
	 */
	GetAccount(ctx context.Context, vo *entity.AccountVo) (*entity.AccountEntity, error)
	/**
	 * @brief: 账户交易
	 * @param: payer
	 * @return: entity.AccountEntity——账户充血模型
	 */
	Transaction(ctx context.Context, payer *entity.AccountEntity, payee *entity.AccountEntity, tx *entity.TransactionVo) (*entity.TransactionEntity, error)
}
