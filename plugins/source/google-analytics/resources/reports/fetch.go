package reports

import (
	"context"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"google.golang.org/api/analyticsreporting/v4"
)

func Fetch(tableName string, request *analyticsreporting.ReportRequest) schema.TableResolver {
	request.HideTotals = true
	request.HideValueRanges = true
	request.PageToken = ""

	return func(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
		c := meta.(*client.Client)
		c.Logger().Info().Msg("started fetch")

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
			request.DateRanges = []*analyticsreporting.DateRange{{StartDate: date, EndDate: date}}

			for {
				resp, err := batchGetReq.Do()
				if err != nil {
					return err
				}

				for _, report := range resp.Reports {
					c.Logger().Info().Bool("golden", report.Data.IsDataGolden).Msg("got report")
					// prt header
					h, err := report.ColumnHeader.MarshalJSON()
					if err != nil {
						return err
					}
					c.Logger().Info().Str("header", string(h)).Msg("got header")

					for _, row := range report.Data.Rows {
						// prt header
						r, err := row.MarshalJSON()
						if err != nil {
							return err
						}
						c.Logger().Info().Str("row", string(r)).Msg("got row")
					}
				}

				report := resp.Reports[0]

				request.PageToken = report.NextPageToken
				if request.PageToken == "" {
					break
				}
			}
		}
		return nil
	}
}
