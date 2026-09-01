package common

const (
	StatusNormal uint8 = 1 // 正常状态
	StatusBanned uint8 = 2 // 禁用状态
)

const (
	DefaultParentID uint64 = 1000000 // 默认父级 ID
	// DefaultInvalidRoleID 用于判断 Token 是否属于 Core
	DefaultInvalidRoleID uint64 = 1000000
)
