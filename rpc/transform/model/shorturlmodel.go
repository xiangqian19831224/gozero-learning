package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ShorturlModel = (*customShorturlModel)(nil)

type (
	// ShorturlModel is an interface to be customized, add more methods here,
	// and implement the added methods in customShorturlModel.
	// ShorturlModel 是一个接口类型，它继承自 shorturlModel 接口
	ShorturlModel interface {
		shorturlModel
	}

	//类型嵌入是一种特殊的语法，允许一个类型“包含”另一个类型。通过类型嵌入，可以实现代码复用和接口组合等功能
	//1. 嵌入方式有两种  直接嵌入类别和嵌入类别的指针
	//2. customShorturlModel 是一个结构体类型，它嵌入了一个 *defaultShorturlModel 类型的字段。
	//3. 通过嵌入，customShorturlModel 可以直接使用 defaultShorturlModel 中定义的所有方法和字段
	customShorturlModel struct {
		*defaultShorturlModel
	}
)

// NewShorturlModel returns a model for the database table.
// customShorturlModel是ShorturlModel接口的一个具体实现，通过嵌入*defaultShorturlModel，
// 它不仅继承了defaultShorturlModel的方法和字段，还满足了ShorturlModel接口的要求，
// 可以在程序中通过ShorturlModel接口来操作customShorturlModel实
func NewShorturlModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ShorturlModel {
	return &customShorturlModel{
		defaultShorturlModel: newShorturlModel(conn, c, opts...),
	}
}
