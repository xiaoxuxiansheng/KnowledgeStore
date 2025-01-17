package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
)

/**
 * @brief: store 仓储服务
 */
type Repository interface {
	/**
	 * @brief: 创建店铺
	 * @param: entity.StoreEntity——店铺充血模型
	 */
	CreateStore(ctx context.Context, entity *entity.StoreEntity) error
	/**
	 * @brief: 获取店铺信息
	 * @param: opts——查询条件
	 * @return: entity.StoreEntity——店铺充血模型
	 */
	GetStores(ctx context.Context, opts ...north.QueryOption) ([]*entity.StoreEntity, error)
}
