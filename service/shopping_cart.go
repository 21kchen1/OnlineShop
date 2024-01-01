package service

/**
 * @File : shopping_cart.go
 * @Description : 购物车管理相关的服务
 * @Author : lei
 * @Date : 2023-12-31
 */

import (
	"onlineshop/models"
)

// GetShopIDByShopName 根据商铺名称获取商铺ID
func GetShopIDByShopName(shopName string) (int, error) {
	// 调用model的方法获取商铺id
	shopID, err := models.GetShopIDByShopName(shopName)
	if err != nil {
		return 0, err
	}

	return shopID, nil
}

// GetProductIDByProductName 根据商品名称获取商品ID
func GetProductIDByProductName(shopName string) (int, error) {
	// 调用model方法获取商品id
	shopID, err := models.GetProductIDByProductName(shopName)
	if err != nil {
		return 0, err
	}

	return shopID, nil
}
