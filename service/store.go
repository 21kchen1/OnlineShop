package service

import (
	"onlineshop/models"
)

// GetStoreIDByUserID 获取商铺ID服务函数
func GetStoreIDByUserID(userID int) (storeID int, err error) {
	// 调用数据库模型的方法获取商铺信息
	store, err := models.GetStoreByUserID(userID)
	if err != nil {
		return 0, err
	}

	return int(store.ID), nil
}

// GetProductsByStoreID 获取商铺下的商品服务函数
func GetProductsByStoreID(storeID int) (products []map[string]interface{}, err error) {
	// 调用数据库模型的方法获取商铺下的商品信息
	storeProducts, err := models.GetProductsByStoreID(storeID)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	for _, product := range storeProducts {
		productData := map[string]interface{}{
			"id":    product.StoreId,
			"name":  product.ProductName,
			"stock": product.Stock,
			// 添加其他需要返回的商品信息字段
		}
		products = append(products, productData)
	}

	return products, nil
}
