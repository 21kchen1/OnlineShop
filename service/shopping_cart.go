package service

/**
 * @File : shopping_cart.go
 * @Description : 购物车管理相关的服务
 * @Author : lei
 * @Date : 2023-12-31
 */

import (
	"errors"
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

// GetShoppingCartDataByUserID 根据用户ID获取购物车数据
func GetShoppingCartDataByUserID(userID int) (productList []map[string]interface{}, err error) {
	// 根据用户ID获取购物车信息
	userLinkShoppingcart, err := models.GetUserLinkShoppingcartByUserID(userID)
	if err != nil {
		return nil, err
	}

	// 根据购物车id获取购物车商品数据
	cartData, err := models.GetShoppingCartLinkProductByShoppingCartID(userLinkShoppingcart.ShoppingCartID)
	if err != nil {
		return nil, err
	}

	for _, item := range cartData {
		product, _ := models.GetProductByID(item.ProductID)
		repro := map[string]interface{} {
			"id": item.ProductID,
			"name": product.ProductName,
			"price": product.Price,
			"quantity": item.Quantity,
			"storeId": product.StoreId,
			"storeName": "黑黑",
		}
		productList = append(productList, repro)
	}
	return
}

// CheckUserShoppingCartLinkExists 检查 user_id - shopping_cart_id 是否存在
func CheckUserShoppingCartLinkExists(userID, shoppingCartID int) (bool, error) {
	return models.CheckUserShoppingCartLinkExists(userID, shoppingCartID)
}

// CheckShoppingCartProductLinkExists 检查 shopping_cart_id - product_id 是否存在
func CheckShoppingCartProductLinkExists(shoppingCartID, productID int) (bool, error) {
	return models.CheckShoppingCartProductLinkExists(shoppingCartID, productID)
}

// UpdateShoppingCartProductQuantity 更新购物车商品项数量
func UpdateShoppingCartProductQuantity(shoppingCartID, productID, quantity int) error {
	exists, err := models.CheckShoppingCartProductLinkExists(shoppingCartID, productID)
	if err != nil {
		return err
	}

	if !exists {
		return errors.New("购物车项不存在")
	}

	return models.UpdateShoppingCartProductQuantity(shoppingCartID, productID, quantity)
}
