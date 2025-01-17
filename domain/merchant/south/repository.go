package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/entity"
)

/**
 * @brief: merchant 仓储服务
 */
type Repository interface {
	/**
	 * @brief: 创建商家
	 * @param: entity.MerchantEntity——商家充血模型
	 */
	CreateMerchant(ctx context.Context, entity *entity.MerchantEntity) error
	/**
	 * @brief: 获取商家信息
	 * @param: merchantId——商家唯一 id
	 * @return: entity.MerchantEntity——商家充血模型
	 */
	GetMerchant(ctx context.Context, merchantId string) (*entity.MerchantEntity, error)
}
