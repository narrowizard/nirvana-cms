package api

import (
	"go-cms/controllers/users"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"
)

// User 用户模块api定义
var User = definition.Descriptor{
	Path:        "/user",
	Description: "user manager api",
	Children: []definition.Descriptor{
		{
			Path:        "/list",
			Description: "get user list",
			Definitions: []definition.Definition{
				{
					Method:   definition.Get,
					Function: users.List,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "page",
							Description: "page index of user list",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
					},
					Results: definition.DataErrorResults("user list data"),
				},
			},
		},
	},
}
