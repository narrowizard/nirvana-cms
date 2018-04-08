package meta

import (
	"github.com/caicloud/nirvana/errors"
)

var (
	TableQueryError  = errors.InternalServerError.Build("Database:TableQueryError", "table ${tableName} query error.")
	UserExistedError = errors.Conflict.Build("Register:UserExistedError", "user ${userName} already existed.")
)
