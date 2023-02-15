package reports

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/analyticsreporting/v4"
)

func fetch(tableName string, request *analyticsreporting.ReportRequest) schema.TableResolver {
	request.HideTotals = true
	request.HideValueRanges = true
	request.PageToken = ""

	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)

		request.ViewId = c.ViewID
		batchGetReq := c.Reports.BatchGet(
			&analyticsreporting.GetReportsRequest{
				ReportRequests: []*analyticsreporting.ReportRequest{request},
			},
		).Context(ctx)

		dates, err := genDates(ctx, c, tableName)
		if err != nil {
			return err
		}

		for date := range dates {
			dateStr := date.Format(dateLayout)
			request.DateRanges = []*analyticsreporting.DateRange{{StartDate: dateStr, EndDate: dateStr}}

			for {
				resp, err := batchGetReq.Do()
				if err != nil {
					return err
				}

				// TODO: parse header

				report := resp.Reports[0]

				for _, row := range report.Data.Rows {
					// parse row
				}

				request.PageToken = report.NextPageToken
				if request.PageToken == "" {
					break
				}
			}
		}
		return nil
	}
}
