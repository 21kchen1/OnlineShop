/**
 * @File : product.go
 * @Description : 商品管理相关的服务
 * @Author : you
 * @Date : 2023-12-18
 */

package service

import (
	"onlineshop/models"
)

// GetProductList 获取商品列表服务函数
func GetProductList(searchKey, productType string) (productList []map[string]interface{}, err error) {
	// 调用数据库模型的方法获取商品列表
	products, err := models.GetProductList(searchKey, productType)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	for _, product := range products {
		productData := map[string]interface{}{
			"productId": product.ID,
			// 添加其他需要返回的商品信息字段
		}
		productList = append(productList, productData)
	}

	return productList, nil
}

/**
 * @File : product.go
 * @Description : 获取商品信息
 * @Author : chen
 * @Date : 2023-12-26
 */
func GetProduct(productId uint) (product map[string]interface{}, err error) {
	theProduct, err := models.GetProductByID(int(productId))

	if err != nil {
		return nil, err
	}

	product = map[string]interface{}{
		"name":      theProduct.ID,
		"price":     theProduct.Price,
		"label":     theProduct.Description,
		"sellerId":  theProduct.StoreId,
		"creatData": theProduct.CreatedAt,
		"sellerNum": theProduct.MonthNum,
	}

	return
}

/**
 * @File : product.go
 * @Description : 根据id删除商品
 * @Author : chen
 * @Date : 2023-12-26
 */
func DeleteProduct(productId int) (err error) {
	err = models.DeleteProductByID(productId)
	return
}