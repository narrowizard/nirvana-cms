package api

import (
	"github.com/narrowizard/nirvana-cms/controllers/users"

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
					Function: users.New,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "account",
							Description: "user account",
							Operators:   []definition.Operator{validator.String("min=5"), validator.String("max=16")},
						},
						{
							Source:      definition.Form,
							Name:        "password",
							Description: "user password",
							Operators:   []definition.Operator{validator.String("min=6"), validator.String("max=20")},
						},
						{
							Source:      definition.Form,
							Name:        "roleid",
							Description: "user role id",
						},
					},
					Results: definition.DataErrorResults("user info"),
				},
			},
		},
		{
			Path:        "/delete",
			Description: "delete user",
			Definitions: []definition.Definition{
				{
					Method:   definition.Delete,
					Function: users.Delete,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "userid",
							Description: "user id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
					},
					Results: []definition.Result{
						{
							Destination: definition.Error,
						},
					},
				},
			},
		},
		{
			Path:        "/update",
			Description: "update user",
			Definitions: []definition.Definition{
				{
					Method:   definition.Update,
					Function: users.Update,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "userid",
							Description: "user id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
						{
							Source:      definition.Form,
							Name:        "status",
							Description: "user status",
						},
						{
							Source:      definition.Form,
							Name:        "roleid",
							Description: "user role id",
						},
					},
					Results: definition.DataErrorResults("whether updated"),
				},
			},
		},
		{
			Path:        "/info",
			Description: "get user info",
			Definitions: []definition.Definition{
				{
					Method:   definition.Get,
					Function: users.Info,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "id",
							Description: "user id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
					},
					Results: definition.DataErrorResults("user info"),
				},
			},
		},
		{
			Path:        "/menulist",
			Description: "get user menu list",
			Definitions: []definition.Definition{
				{
					Method:   definition.List,
					Function: users.MenuList,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "nirvanacmsuserid",
							Description: "user id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
					},
					Results: definition.DataErrorResults("user menu list data"),
				},
			},
		},
	},
}
