# 平台简介

​		本平台基于vue、gin、gorm框架实现，服务于各个年龄段和不同需求的消费者群体，为各种不同出售类型的商家群体提供售卖平台。

 

## 平台功能

​    	本系统为ECS（electronic commerce system）购物平台，主要为消费者与商家提供线上交易途径。消费者方面，平台提供订单、购物车、收藏夹与搜索等功能，为消费者提供高效、便利的消费体验。商家方面，在消费者功能的基础上，平台还将额外提供商品管理和订单管理功能，为商家提供直接、快速的售卖途径。管理员方面，为方便平台管理，在消费者功能的基础上，平台还将额外提供用户管理、开发日志等功能，提供强大的平台管理能力。

* 对于消费者。
  * 在购买方面，提供商品搜索，商品查看，订单的商品的增删查改，收藏夹的商品的增删查改，购物车的商品的增删查改。
  * 消费者可以创建多个收藏夹，实现对商品的分类收藏。在售后方面，提供消费者对商品、商家评价。

* 对于商家。
  * 在消费者功能的基础上，提供额外功能。
  * 在商品出售方面，提供商品增删查改，商品订单的获取和商品的发货功能。

* 对于管理员。
  * 在消费者功能的基础上，提供额外功能。
  * 在用户管理方面，提供用户类型（权限）修改，用户的增删查改。
  * 在评论管理方面，提供评论的增删查改。
  * 在开发记录方面，提供开发日志的增删改。

  	此外，平台提供基本的账号注册，登录功能。



## 平台使用

>前端

本平台前端使用 vue，bootstrap 框架实现，启动前端服务前需要安装如下内容：

* node.js
* 相关依赖安装

```html
VueRouter
Vuex
npm install bootstrap-vue
npm install bootstrap@5.3.0-alpha1
npm i element-ui -S
npm install axios --save
npm install vue-cookies --save
npm install gsap 
npm install echarts
```

* 启动前端服务

```html
npm run serve
```

* 在 localhost:8080 进入服务页面

>后端

本平台后端使用 go 结合 gin、gorm 实现，启动后端服务前需要安装如下内容：

* golang
* gin 安装

```
go get -u github.com/gin-gonic/gin
```

* gorm 安装

```
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

* 启动后端服务

从代码文件中的 main.go 文件启动