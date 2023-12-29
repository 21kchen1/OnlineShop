# **前后端API接口规范**

绿色标记为测试完成的接口，不允许更改

黄色标记为代码编写完成的接口，不允许更改

删除标记为删除的接口

红色标记为有问题的接口



1、     登入接口

url：/user/login

请求类型：POST

请求参数：username,password

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示),userId(若成功返回用户ID)



2、     注册接口

url：/user/register

请求类型：POST

请求参数：username,password,email

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



3、     获取用户信息

url：/user/getInf

请求类型：POST

请求参数：userId

返回参数：username,email,p



4、     获取商品列表-

url：/product/getList

请求类型：POST

请求参数：searchKey(搜索关键词),type(商品类型),

返回参数：data:[ { productId(商品ID) } , { } , ...]



5、     获取商品信息

url：/product/getProductInfo

请求类型：POST

请求参数：productId

返回参数：productId, productName, description, storeId, price, stock, status



6、     添加商品-

url：/product/addProduct

请求类型：POST

请求参数：storeId(店铺id)，productName(商品名称)，description（商品描述），stock（商品数量），price（商品价格）

返回参数：isDelete(boolen类型),msg(成功或错误信息提示)



7、     商品删除

url：/product/deleteProduct

请求类型：POST

请求参数：productId

返回参数：isDelete(boolen类型),msg(成功或错误信息提示)



8、     修改商品信息

url：/product/editProduct

请求类型：POST

请求参数：productId, productName, description, storeId, price, stock, productStatus

返回参数：isDelete(boolen类型),msg(成功或错误信息提示)



9、     获取商品库存

url：/product/getQuantity

请求类型：POST

请求参数：productId

返回参数：quantity(int类型)



10、     修改商品库存

url：/product/editQuantity

请求类型：POST

请求参数：productId,editQuantity(修改后的库存量)

返回参数：isSuccess(boolen类型)



11、     减少商品库存

url：/product/editQuantity

请求类型：POST

请求参数：productId,editQuantity(修改后的库存量)

返回参数：isSuccess(boolen类型)



14、     获取用户订单列表

url：/order/useGetList

请求类型：POST

请求参数：userId

返回参数：

data:[{

  productId(商品ID)

  productName(商品名称)

  orderNumber(订单号),

  price（商品价格）

  deliveryStatus（发货状态）

  date(下单的时间）

}]



15、     获取用户收藏夹列表

url：/favorites/geList

请求类型：POST

请求参数：userId(用户ID)

返回参数：data:[{ id(收藏夹ID,name(收藏夹名字),num(收藏商品的数量)},{}..] 



16、     获取收藏夹中商品列表

url：/favorites/getProduct

请求类型：POST

请求参数：favoritesId(收藏夹ID)

返回参数：data:[ { productId(商品ID) }, { } , .....]



17、     用户创建收藏夹

url：/favorites/add

请求类型：POST

请求参数：userId,favoritesName(收藏夹名称)

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



18、     用户删除收藏夹

url：/favorites/delete

请求类型：POST

请求参数：favoritesId(收藏夹ID)

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



19、     用户修改收藏夹名称

url：/favorites/updata

请求类型：POSTN

请求参数：favoritesId(收藏夹ID),favoritesName(收藏夹名称)

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



20、     收藏夹中添加商品

url：/favorites/addProduct

请求类型：POST

请求参数：favoritesId,productId

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



21、     收藏夹中删除商品

url：/favorites/deleteProduct

请求类型：POST

请求参数：favoritesId,productId

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



22、     获取用户浏览历史

url：/history/getBrowse

请求类型：POST

请求参数：productId，productImg，productTitle，productPrice

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



23、     用户清空浏览历史

url：/history/delete

请求类型：POST

请求参数：productId

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



 27、     获取商品评论

url：/product/getComment

请求类型：POST

请求参数：productId(商品ID)

返回参数：(当参数特别多的时候要用JSON代码块)



```
data:[{ 
    commentId(评论ID),
    userId(发表者ID),
    username(发表者用户名),
    content(评论内容),
    time(发表时间),
    goodNum(点赞数),
    star(评价星级),
    reply(回复):[{
        commentId(评论ID),
        userId(发表者ID),
        username(发表者用户名),
        content(评论内容),
        time(发表时间),
        goodNum(点赞数),
    }]
}]
```



28、     用户发表评论

url：/comment/add

请求类型：POST

请求参数：userId,productId,content,star(星级)

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



29、     用户回复评论

url：/comment/reply

请求类型：POST

请求参数：userId,commentId,content

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



30、     修改用户信息

url：/user/updata

请求类型：POST

请求参数：editedUser{id,username,email,phone,password,role}

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(用户ID)，

  username(用户名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 role（角色即权限）

}]



31、     添加用户

url：/user/addUser

请求类型：POST

请求参数：newsUser{username、email、phone、password、role}

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(用户ID)，

  username(用户名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 role（角色即权限）

}]



32、     获取用户列表

url:/user/getList

请求类型：GET

请求参数：返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(用户ID)，

  username(用户名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 role（角色即权限）

}]



33、     删除用户

url：/user/delete

请求类型：POST

请求参数：userID

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(用户ID)，

  username(用户名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 role（角色即权限）

}]



34、     获取商家列表

url:/seller/getList

请求类型：GET

请求参数：返回参数：

data:[{

  id(商家ID)，

  sellername(商家名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 address（商家地址）

}]



35、     添加商家

url：/sellerr/addSeller

请求类型：POST

请求参数：newSeller{sellername、email、phone、address、password}

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(商家ID)，

  sellername(商家名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 password,

​	 address（商家地址）

}]



36、     修改商家信息

url：/seller/updata

请求类型：POST

请求参数：editedSeller

{sellerId,sellername,email,phone,address,password}

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  id(商家ID)，

  sellername(商家名称)，

​	 mail（邮箱），

​	 phone（号码），

​	 password,

​	 address（商家地址）

}]



37、     删除商家

url：/seller/delete

请求类型：POST

请求参数：sellerId

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



38、     获取所有订单列表

url：/oder/orderList

请求类型：POST

请求参数：

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)

data:[{

  Id(数字)

  Name(订单用户)

  address(订单地址),

  title(日志的标题, 需要字符串),

  orderNumber(订单号),

  price（商品价格）

  productImg（商品图片）

  deliveryStatus（发货状态）

  date(下单的时间）

}]



39、     修改订单信息

url：/order/editOrder

请求类型：POST

请求参数：orderId，orderNumber（订单号），orderAdress（订单地址），deliveryStatus（订单发货状态）

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



40、     删除订单

url：/order/deleteOrder

请求类型：POST

请求参数：orderId(订单数据的id）

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



41、     删除评论

url：/comment/delete

请求类型：POST

请求参数：commentId（这个评论数据的ID）

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



 42、     获取日志列表

url：/log/getInf

请求类型：POST

请求参数：

返回参数：(当参数特别多的时候要用JSON代码块)

```
data:[{
    Id(数字)
    submitName(提交的人)
    time_data(提交时间，需要“年/月/日 时:分:秒”的格式，是字符串),
    title(日志的标题, 需要字符串),
    text(日志的内容，需要字符串),
}]
```



43、     添加日志

url：/log/addLog

请求类型：POST

请求参数：userId，title，content

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



44、     删除日志

url：/log/delLog

请求类型：POST

请求参数：logId

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



45、     修改日志

url：/log/edit

请求类型：POST

请求参数：logId，title, content

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



vue发送post请求实例

```
import { post } from "@/utils/http";

export default {
  mounted() {
  //""填入url地址，{}为请求参数
    post("", {}).then(
      (Response) => {
        console.log("请求成功", Response);
        //Response是返回的参数
      },
      (error) => {
        console.log("请求失败", error.message);
      }
    );
  },
};
```



46、     发布商品

url：/api/products

请求类型：POST

请求参数：productName：商品名称，productDescription：商品描述，productPrice：商品价格，stockQuantity：库存数量

返回参数：isSuccess(boolen类型),msg(成功或错误信息提示)



47、获取商铺id

url：/api/getShopInfo

请求类型：POST

请求参数：userId

返回参数：shopid



48、根据商铺id获取商品

url：/api/getShopProducts

请求类型：POST

请求参数：shopId

返回参数：id（商品id），name（商品名称），stock（库存数量）



49、更新商品库存

url：/api/updateProductStock

请求类型：POST

请求参数：shopId（商铺id），productId（商品id），newStock（新的库存数量）

返回参数：isSuccess(boolen类型)



购物车的四个接口：

50、获取当前账号购物车数据

url：/api/cart

请求类型：GET

请求参数：Userid

返回参数：以下格式数据实例，selected默认false即可



51、根据商铺名称获取商铺id

url：/api/shopId

请求类型：GET

请求参数：shopName（商铺名字）

返回参数：shopId



52、根据商品名称获取商品id

url：/api/productId

请求类型：GET

请求参数：productName(商品名字)

返回参数：productId



53、

url：/api/updataQuantity

请求类型：POST

请求参数：userId,shopId,productId,quantity

返回参数：返回参数：isSuccess(boolen类型)









​            54、     主页获取12个推荐商品（随机）

url：/product/recommend

请求类型：POST

请求参数：

返回参数：data:[ { productId(商品id) } ]