package client

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

var ViewIDColumn = schema.Column{
	Name:        "view_id",
	Type:        schema.TypeString,
	Description: "View ID",
	Resolver:    ResolveViewID,
	CreationOptions: schema.ColumnCreationOptions{
		PrimaryKey: true,
		NotNull:    true,
	},
}
