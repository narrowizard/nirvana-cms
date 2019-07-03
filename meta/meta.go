package meta

import (
	"github.com/caicloud/nirvana/errors"
)

var (
	TableQueryError      = errors.InternalServerError.Build("Database:TableQueryError", "table ${tableName} query error.")
	TableInsertError     = errors.InternalServerError.Build("Database:TableInsertError", "table ${tableName} insert error.")
	TableUpdateError     = errors.InternalServerError.Build("Database:TableUpdateError", "table ${tableName} update error.")
	UserExistedError     = errors.Conflict.Build("Register:UserExistedError", "user ${userName} already existed.")
	UnexpectedParamError = errors.BadRequest.Build("Params:UnexpectedParamError", "unexpected param: ${paramName}.")
	RoleExistedError     = errors.Conflict.Build("Role:RoleExistedError", "role ${roleName} already existed.")
)
