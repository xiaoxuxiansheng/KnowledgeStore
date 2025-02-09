package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
)

/**
 * @brief: store domain 服务
 */
type IService interface {
	/**
	 * @brief: 创建店铺
	 * @param: entity.StoreVo——店铺 vo
	 * @return: entity.StoreEntity——店铺充血模型
	 */
	CreateStore(ctx context.Context, req *entity.StoreEntity) error
	OnlineStore(ctx context.Context, req *entity.StoreEntity) error
	OfflineStore(ctx context.Context, req *entity.StoreEntity) error

	/**
	 * @brief: 查询店铺
	 * @param: opts——查询条件
	 * @return: entity.StoreEntity——店铺充血模型
	 */
	GetStores(ctx context.Context, opts ...QueryOption) ([]*entity.StoreEntity, int, error)
}
