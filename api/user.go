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
					Method:   definition.List,
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
						{
							Source:      definition.Query,
							Name:        "pagesize",
							Description: "pagesize of user list",
							Operators:   []definition.Operator{validator.Int("min=1"), validator.Int("max=20")},
						},
						{
							Source:      definition.Query,
							Name:        "search",
							Description: "search info",
						},
					},
					Results: definition.DataErrorResults("user list data"),
				},
			},
		},
		{
			Path:        "/new",
			Description: "create user",
			Definitions: []definition.Definition{
				{
					Method:   definition.Create,
					Function: users.List,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Body,
							Name:        "account",
							Description: "user account",
							Operators:   []definition.Operator{validator.String("min=5"), validator.String("max=16")},
						},
						{
							Source:      definition.Body,
							Name:        "password",
							Description: "user password",
							Operators:   []definition.Operator{validator.String("min=6"), validator.String("max=20")},
						},
					},
					Results: definition.DataErrorResults("whether created"),
				},
			},
		},
	},
}
