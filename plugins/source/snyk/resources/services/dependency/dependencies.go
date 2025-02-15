package dependency

import (
	"context"
	"fmt"

	"github.com/cloudquery/cloudquery/plugins/source/snyk/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"github.com/pavel-snyk/snyk-sdk-go/snyk"
)

func Dependencies() *schema.Table {
	return &schema.Table{
		Name:        "snyk_dependencies",
		Description: `https://snyk.docs.apiary.io/#reference/dependencies/dependencies-by-organization/list-all-dependencies`,
		Resolver:    fetchDependencies,
		Multiplex:   client.ByOrganization,
		Transform:   transformers.TransformWithStruct(&snyk.Dependency{}, transformers.WithPrimaryKeys("ID")),
		Columns:     schema.ColumnList{client.OrganizationID},
	}
}

func fetchDependencies(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	result, _, err := c.Dependencies.List(ctx, c.OrganizationID)
	if err != nil {
		return err
	}
	for _, dep := range result {
		res <- dep
		fmt.Println(dep)
	}

	res <- result

	return nil
}
