// Package cache Cache 接口
package cache

import "context"

// Cache 利用 go 接口的特性，将所有的缓存操作都抽象到这里 实现与存储无关的缓存操作
type Cache interface {
	// IsEmpty 判断缓存是否为空
	IsEmpty(value string) bool
	// Get 获取缓存
	Get(ctx context.Context, key string) (string, error)
	// MGet 批量获取缓存
	MGet(ctx context.Context, keys ...string) (map[string]string, error)
	// PrefixGet 前置匹配获取
	PrefixGet(ctx context.Context, prefix string, scanCount int64) (map[string]string, error)
	// Set 设置缓存
	Set(ctx context.Context, key string, value string) error
	// MSet 批量设置缓存
	MSet(ctx context.Context, data map[string]string) error
	// Del 删除缓存
	Del(ctx context.Context, key string) error
	// MDel 批量删除缓存
	MDel(ctx context.Context, keys ...string) error
	// 前缀匹配删除
	PrefixDel(ctx context.Context, prefix string, scanCount int64) error
	// Clear 清空缓存
	Clear(ctx context.Context) error
	// Close 关闭缓存
	Close() error
	// 获取缓存类型
	GetType() string
}
