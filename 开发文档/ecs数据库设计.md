# 用户类

* User（用户）
  * 用户编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 手机电话 PhoneNum
  * 用户类型 UserType
  * 邮箱 Email



* Seller（商家）
  * 商家编号 ID（key）
  * 用户编号 UserID（foreignKey User）
  * 地址 Address



* Store（店铺）
  * 店铺编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 用户编号 UserID（foreignKey User）
  * 店铺名称 StoreName
  * 联系信息 ContactInfo
  * 关注数量 Followers
  * 产品数量 ProductCount
  * 销售数量 SalesNum



# 资源类

* Product（产品）
  * 产品编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 店铺编号 StoreId（foreignKey Store）
  * 产品名称 ProductName 
  * 描述 Description 
  * 产品状态 ProductStatus
  * 月销量 MonthNum
  * 库存 Stock
  * 产品类型 ProductType
  * 订约数 Likes
  * 评论数 Comments
  * 价格 Price



* Comment（评论）
  * 评论编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 用户编号 UserID（foreignKey User）
  * 产品编号 ProductID（foreignKey Product）
  * 点赞数 Likes
  * 评论内容 Content



# 表单类

* Order（订单）
  * 订单编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 产品编号 ProductID（foreignKey Product）
  * 订单状态 OrderStatus
  * 总价 TotalPrice
  * 数量 Quantity
  * 下单时间 OrderTime
  * 支付时间 PayTime
  * 发货时间 ShippingTime
  * 完成时间 CompletionTime



* Log（日志）
  * 日志编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 用户编号 ID（foreignKey User）
  * 标题 Title
  * 内容 Content



* Favorites（收藏夹）
  * 收藏夹编号 ID（key）
  * 创建时间 CreatedAt
  * 更新时间 UpdatedAt
  * 删除时间 DeletedAt
  * 用户编号 ID（foreignKey User）
  * 收藏夹名称 FavoritesName
  * 收藏数量 Count



* FavoritesLinkProduct（收藏夹产品链接）
  * 收藏夹产品链接编号 ID（key）
  * 收藏夹编号 FavoritesID（foreignKey  Favorites）
  * 产品编号 ProductID（foreignKey  Product）