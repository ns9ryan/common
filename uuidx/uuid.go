package uuidx

import (
	"github.com/gofrs/uuid/v5"
	"github.com/zeromicro/go-zero/core/logx"
)

// NewUUID 生成 UUID
func NewUUID() uuid.UUID {
	id, err := uuid.NewV7()
	if err != nil {
		logx.Errorw("生成 UUID 失败", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return id
}

// ParseUUIDString 将 UUID 字符串转换为 UUID
func ParseUUIDString(id string) uuid.UUID {
	result, err := uuid.FromString(id)
	if err != nil {
		logx.Errorw("解析 UUID 失败", logx.Field("detail", err))
		return uuid.UUID{}
	}
	return result
}

// ParseUUIDSlice 将 UUID 字符串切片转换为 UUID 切片
func ParseUUIDSlice(ids []string) []uuid.UUID {
	result := make([]uuid.UUID, 0, len(ids))

	for _, v := range ids {
		p, err := uuid.FromString(v)
		if err != nil {
			logx.Errorw("解析 UUID 失败", logx.Field("detail", err))
			return nil
		}
		result = append(result, p)
	}
	return result
}

// ParseUUIDStringToPointer 将 UUID 字符串指针转换为 UUID 指针
func ParseUUIDStringToPointer(id *string) *uuid.UUID {
	if id == nil {
		return nil
	}

	result, err := uuid.FromString(*id)
	if err != nil {
		logx.Errorw("解析 UUID 失败", logx.Field("detail", err))
		return nil
	}
	return &result
}

// ParseUUIDSliceToPointer 将 UUID 字符串切片转换为 UUID 指针切片
func ParseUUIDSliceToPointer(ids []string) []*uuid.UUID {
	result := make([]*uuid.UUID, 0, len(ids))

	for _, v := range ids {
		p, err := uuid.FromString(v)
		if err != nil {
			logx.Errorw("解析 UUID 失败", logx.Field("detail", err))
			return nil
		}
		result = append(result, &p)
	}
	return result
}
