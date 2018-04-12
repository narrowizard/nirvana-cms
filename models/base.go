package models

import (
	"time"
)

// ENUMSTATUS 状态
type ENUMSTATUS int

const (
	// STATUSNORMAL 正常状态
	STATUSNORMAL ENUMSTATUS = 1
	// STATUSBLOCK 禁用状态
	STATUSBLOCK ENUMSTATUS = 2
	// STATUSDELETED 已删除
	STATUSDELETED ENUMSTATUS = 101
)

// Model db base model
type Model struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

// PagableData 分页数据
type PagableData struct {
	Total int
	Data  interface{}
}
