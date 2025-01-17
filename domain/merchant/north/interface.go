package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/entity"
)

/**
 * @brief: merchant domain 服务
 */
type IService interface {
	/**
	 * @brief: 注册用户
	 * @return: entity.MerchantEntity——商家充血模型
	 */
	Register(ctx context.Context) (*entity.MerchantEntity, error)
	/**
	 * @brief: 查询商家
	 * @param: merchantId——商家唯一 id
	 * @return: entity.MerchantEntity——商家充血模型
	 */
	GetMerchant(ctx context.Context, merchantId string) (*entity.MerchantEntity, error)
}
