package api

import (
	"go-cms/controllers/menus"

	"github.com/caicloud/nirvana/definition"
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
							Name:        "parentid",
							Description: "parent node's id",
						},
					},
					Results: definition.DataErrorResults("whether created"),
				},
			},
		},
	},
}
