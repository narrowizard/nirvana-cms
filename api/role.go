package api

import (
	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"
	"github.com/narrowizard/nirvana-cms/controllers/roles"
)

// Role 角色模块api定义
var Role = definition.Descriptor{
	Path:        "/role",
	Description: "role manager api",
	Children: []definition.Descriptor{
		{
			Path:        "/list",
			Description: "list role",
			Definitions: []definition.Definition{
				{
					Method:   definition.List,
					Function: roles.List,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "search",
							Description: "search keyword",
						},
					},
					Results: definition.DataErrorResults("role list data"),
				},
			},
		},
		{
			Path:        "/info",
			Description: "get role info",
			Definitions: []definition.Definition{
				{
					Method:   definition.Get,
					Function: roles.Info,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "roleid",
							Description: "role id",
						},
					},
					Results: definition.DataErrorResults("role info"),
				},
			},
		},
		{
			Path:        "/new",
			Description: "create role",
			Definitions: []definition.Definition{
				{
					Method:   definition.Create,
					Function: roles.New,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "name",
							Description: "role name",
						},
						{
							Source:      definition.Form,
							Name:        "menus",
							Description: "user menu list",
						},
					},
					Results: definition.DataErrorResults("role info"),
				},
			},
		},
		{
			Path:        "/delete",
			Description: "delete role",
			Definitions: []definition.Definition{
				{
					Method:   definition.Delete,
					Function: roles.Delete,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Query,
							Name:        "roleid",
							Description: "role id",
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
			Description: "update role",
			Definitions: []definition.Definition{
				{
					Method:   definition.Update,
					Function: roles.Update,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "roleid",
							Description: "role id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
						{
							Source:      definition.Form,
							Name:        "name",
							Description: "role name",
						},
						{
							Source:      definition.Form,
							Name:        "menus",
							Description: "role menu list",
						},
					},
					Results: definition.DataErrorResults("whether updated"),
				},
			},
		},
	},
}
