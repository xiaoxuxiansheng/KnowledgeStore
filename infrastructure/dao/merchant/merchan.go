package merchant

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/south"
)

type Dao struct{}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateMerchant(ctx context.Context, entity *entity.MerchantEntity) error {
	return nil
}

/**
 * @brief: 获取商家信息
 * @param: merchantId——商家唯一 id
 * @return: entity.MerchantEntity——商家充血模型
 */
func (d *Dao) GetMerchant(ctx context.Context, merchantId string) (*entity.MerchantEntity, error) {
	return nil, nil
}
