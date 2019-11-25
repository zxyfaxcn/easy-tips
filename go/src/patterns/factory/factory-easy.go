package main

//---------------
//叼叼的设计模式(Go享版)
//简单工厂模式
//@auhtor TIGERB<https://github.com/TIGERB>
//---------------

import "fmt"

const (
	// CartConst 购物车列表页面
	CartConst = "cart/list"
	// ProductConst 商品列表页面
	ProductConst = "product/list"
)

// Context 请求上下文
type Context struct {
	URI string
}

// PageInterface PageInterface
type PageInterface interface {
	MakeData(c *Context) (interface{}, error)
}

// Cart 购物车页面数据对象
type Cart struct {
	header interface{}
	footer interface{}
}

// MakeData 构建数据对象
func (Cart *Cart) MakeData(c *Context) (interface{}, error) {
	// 构建数据的页面代码...
	fmt.Println("生成购物车静态页面数据对象...")
	return Cart, nil
}

// Product Spu页面数据对象
type Product struct {
	header interface{}
	footer interface{}
}

// MakeData 构建数据对象
func (Product *Product) MakeData(c *Context) (interface{}, error) {
	// 构建数据的页面代码...
	fmt.Println("生成spu详情静态页面数据对象...")
	return Product, nil
}

// PageFactory 构建页面对象的简单工厂
type PageFactory struct {
	Ctx *Context
}

// Get 获取对象
func (p *PageFactory) Get() PageInterface {
	switch p.Ctx.URI {
	case CartConst:
		return &Cart{}
	case ProductConst:
		return &Product{}

	default:
		panic("不支持的页面")
	}
}

// 客户端使用示例
func main() {
	ctx := &Context{
		URI: "cart/list",
	}
	pageFactory := &PageFactory{
		Ctx: ctx,
	}
	pageFactory.Get().MakeData(ctx)
}
