package utils

import (
	"cmp"
	"slices"
)

// SliceSortCompact 切片去重，Go语言1.21 slices包语法
func SliceSortCompact[T cmp.Ordered](s []T) []T {
	//先排序
	slices.Sort(s)
	//压缩，去重
	return slices.Compact(s)
}
