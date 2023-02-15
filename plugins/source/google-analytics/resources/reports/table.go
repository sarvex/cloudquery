package reports

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/analyticsreporting/v4"
)

func fetchTable(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)

	from, err := getStartDate(ctx, c, "TODO")
	if err != nil {
		return err
	}

	req := &analyticsreporting.ReportRequest{
		DateRanges: []*analyticsreporting.DateRange{
			// TODO: day-by-day increments
			{EndDate: currDate(), StartDate: from},
		},
		DimensionFilterClauses: nil,
		Dimensions:             nil,
		FiltersExpression:      "",
		MetricFilterClauses:    nil,
		Metrics:                nil,
		OrderBys:               nil,
		ViewId:                 c.ViewID,
	}
	batchGetReq := c.Reports.BatchGet(&analyticsreporting.GetReportsRequest{
		ReportRequests: []*analyticsreporting.ReportRequest{req},
	}).Context(ctx)

	for {
		resp, err := batchGetReq.Do()
		if err != nil {
			return err
		}
		report := resp.Reports[0].ColumnHeader.MetricHeader.MetricHeaderEntries[0].Type

		for _, row := range report.Data.Rows {
			// parse row
		}
		req.PageToken = report.NextPageToken
		if req.PageToken == "" {
			return nil
		}
	}
}
