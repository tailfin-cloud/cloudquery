package deployment

import (
	"github.com/cloudquery/cloudquery/plugins/source/vercel/client"
	"github.com/cloudquery/cloudquery/plugins/source/vercel/internal/vercel"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DeploymentChecks() *schema.Table {
	return &schema.Table{
		Name:          "vercel_deployment_checks",
		Resolver:      fetchDeploymentChecks,
		Transform:     client.TransformWithStruct(&vercel.DeploymentCheck{}),
		Multiplex:     client.TeamMultiplex,
		IsIncremental: true,
		Columns: []schema.Column{
			{
				Name:     "deployment_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("uid"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
