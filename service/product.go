/**
 * @File : product.go
 * @Description : 商品管理相关的服务
 * @Author : you
 * @Date : 2023-12-18
 */

package service

import (
	"errors"
	"math/rand"
	"onlineshop/models"
	"reflect"
	"time"
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
 * @Description : 添加商品
 * @Author : chen
 * @Date : 2023-12-26
 */
func AddProduct(theProduct *models.Product) (err error) {
	if theProduct.ProductName == "" {
		err = errors.New("产品名不可为空")
		return
	}
	err = models.CreateAProduct(theProduct)

	return
}

/**
 * @File : product.go
 * @Description : 获取商品信息
 * @Author : chen
 * @Date : 2023-12-26
 */
func GetProduct(productId int) (product map[string]interface{}, err error) {
	theProduct, err := models.GetProductByID(productId)

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

/**
 * @File : product.go
 * @Description : 根据id修改商品信息
 * @Author : chen
 * @Date : 2023-12-27
 */
func EditProduct(productId int, theProduct models.Product) (err error) {
	preProduct, err := models.GetProductByID(productId)

	if err != nil {
		return err
	}

	// 只选择非空的属性进行更新
	typeOfProduct := reflect.TypeOf(&preProduct).Elem()
	valueOfPre := reflect.ValueOf(&preProduct).Elem()
	valueOfNew := reflect.ValueOf(theProduct)

	for i := 0; i < typeOfProduct.NumField(); i++ {
		field := typeOfProduct.Field(i)
		fieldNew := valueOfNew.Field(i)

		// 非空
		if !fieldNew.IsZero() {
			valueOfPre.FieldByName(field.Name).Set(fieldNew)
		}
	}

	err = models.UpdateAProduct(&preProduct)

	return
}

/**
 * @File : product.go
 * @Description : 根据id获取商品数量
 * @Author : chen
 * @Date : 2023-12-26
 */
func GetProductNum(productId int) (num int, err error) {
	theProduct, err := models.GetProductByID(productId)

	if err != nil {
		return -1, err
	}

	return theProduct.Stock, err
}

/**
 * @File : product.go
 * @Description : 根据修改商品数量
 * @Author : chen
 * @Date : 2023-12-26
 */
func EditProductNum(productId int, editQuantity int) (err error) {
	if editQuantity < 0 {
		err = errors.New("商品数量不可小于0")
		return
	}
	theProduct, err := models.GetProductByID(productId)

	if err != nil {
		return err
	}

	theProduct.Stock = editQuantity

	err = models.UpdateAProduct(&theProduct)

	return
}

// GetRecommendedProducts 获取主页推荐商品服务函数
func GetRecommendedProducts() ([]map[string]interface{}, error) {
	// 调用数据库模型的方法获取所有商品
	allProducts, err := models.GetAllProduct()
	if err != nil {
		return nil, err
	}

	// 如果商品数量小于等于12，直接返回所有商品
	if len(allProducts) <= 12 {
		return allProductsToData(allProducts), nil
	}

	// 随机获取12个商品
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(allProducts), func(i, j int) {
		allProducts[i], allProducts[j] = allProducts[j], allProducts[i]
	})

	return allProductsToData(allProducts[:12]), nil
}

// 将商品切片转为前端需要的数据格式
func allProductsToData(products []*models.Product) []map[string]interface{} {
	var productList []map[string]interface{}
	for _, product := range products {
		productData := map[string]interface{}{
			"productId": product.ID,
			// 添加其他需要返回的商品信息字段
		}
		productList = append(productList, productData)
	}
	return productList
}
