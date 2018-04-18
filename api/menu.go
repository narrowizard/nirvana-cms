package api

import (
	"nirvana-cms/controllers/menus"

	"github.com/caicloud/nirvana/definition"
	"github.com/caicloud/nirvana/operators/validator"
)

var Menu = definition.Descriptor{
	Path:        "/menu",
	Description: "menu manager api",
	Children: []definition.Descriptor{
		{
			Path:        "/list",
			Description: "get menu tree data",
			Definitions: []definition.Definition{
				{
					Method:   definition.List,
					Function: menus.List,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Results:  definition.DataErrorResults("menu list data"),
				},
			},
		},
		{
			Path:        "/new",
			Description: "create a menu node",
			Definitions: []definition.Definition{
				{
					Method:   definition.Create,
					Function: menus.New,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "name",
							Description: "menu name",
						},
						{
							Source:      definition.Form,
							Name:        "icon",
							Description: "node icon",
						},
						{
							Source:      definition.Form,
							Name:        "url",
							Description: "node url",
						},
						{
							Source:      definition.Form,
							Name:        "remarks",
							Description: "node remarks",
						},
						{
							Source:      definition.Form,
							Name:        "parentid",
							Description: "parent node's id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
						{
							Source:      definition.Form,
							Name:        "ismenu",
							Description: "1 for menu, 2 for api",
							Operators:   []definition.Operator{validator.Int("min=1"), validator.Int("max=2")},
						},
					},
					Results: definition.DataErrorResults("whether created"),
				},
			},
		},
		{
			Path:        "/update",
			Description: "update a menu node",
			Definitions: []definition.Definition{
				{
					Method:   definition.Update,
					Function: menus.Update,
					Consumes: []string{definition.MIMEAll},
					Produces: []string{definition.MIMEJSON},
					Parameters: []definition.Parameter{
						{
							Source:      definition.Form,
							Name:        "name",
							Description: "menu name",
						},
						{
							Source:      definition.Form,
							Name:        "icon",
							Description: "node icon",
						},
						{
							Source:      definition.Form,
							Name:        "url",
							Description: "node url",
						},
						{
							Source:      definition.Form,
							Name:        "remarks",
							Description: "node remarks",
						},
						{
							Source:      definition.Form,
							Name:        "id",
							Description: "node id",
							Operators:   []definition.Operator{validator.Int("min=1")},
						},
						{
							Source:      definition.Form,
							Name:        "ismenu",
							Description: "1-menu 2-api",
							Operators:   []definition.Operator{validator.Int("min=1"), validator.Int("max=2")},
						},
					},
					Results: definition.DataErrorResults("whether updated"),
				},
			},
		},
	},
}
