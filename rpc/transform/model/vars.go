package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

// 声明了一个名为 ErrNotFound 的错误变量，并将其初始化为 sqlx 包中的 ErrNotFound 错误。
// ErrNotFound 通常用于表示在数据库查询中未找到期望的数据。
var ErrNotFound = sqlx.ErrNotFound
