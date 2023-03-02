package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/resources/reports"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/analyticsreporting/v4"
)

func tables() schema.Tables {
	return schema.Tables{
		{
			Name: "google_analytics_custom",
			Columns: schema.ColumnList{
				{
					Name: "custom_field",
					Type: schema.TypeInt,
				},
			},
			Resolver: reports.Fetch("google_analytics_custom", &analyticsreporting.ReportRequest{
				Dimensions: []*analyticsreporting.Dimension{ // = group by, at most 10
					{Name: "ga:date"},
					{Name: "ga:language"},
					{Name: "ga:country"},
					{Name: "ga:city"},
					{Name: "ga:browser"},
					{Name: "ga:operatingSystem"},
					{Name: "ga:year"},
					{Name: "ga:month"},
					{Name: "ga:hour"},
				},
				Metrics: []*analyticsreporting.Metric{ // required values
					{Expression: "ga:users"},
					{Expression: "ga:newUsers"},
					{Expression: "ga:sessions"},
					{Expression: "ga:sessionsPerUser"},
					{Expression: "ga:pageviews"},
					{Expression: "ga:pageviewsPerSession"},
					{Expression: "ga:avgSessionDuration"},
					{Expression: "ga:bounceRate"},
				},
			}),
			IsIncremental: true,
		},
	}
}
