package reports

import (
	"context"
	"fmt"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
)

// genDates will produce dates in 2006-01-02 layout (YYYY-MM-DD)
func genDates(ctx context.Context, c *client.Client, table string) (<-chan string, error) {
	res, err := c.Backend.Get(ctx, table, c.ID())
	if err != nil {
		return nil, err
	}

	startDate := c.StartDate
	if len(res) > 0 {
		startDate = res
	}

	// parse
	const dateLayout = "2006-01-02"
	t, err := time.Parse(dateLayout, startDate)
	if err != nil {
		return nil, err
	}

	ch := make(chan string)
	defer close(ch)

	today := time.Now().UTC()

	for !t.After(today) {
		ch <- t.Format(dateLayout)
		t.Add(24 * time.Hour)
	}

	return ch, nil
}

// https://ga-dev-tools.google/dimensions-metrics-explorer/time
// Other values aren't supported
var timeFormats = map[string]string{
	"ga:date":           "20060102",
	"ga:year":           "2006",
	"ga:month":          "01",
	"ga:hour":           "03",
	"ga:minute":         "04",
	"ga:dateHour":       "2006010203",
	"ga:dateHourMinute": "200601020304",
}

// TODO: make sure the zone is returned correctly?
// https://groups.google.com/g/google-analytics-data-export-api/c/4A0-qoyRuOU
func parseTime(field string, value string) (*time.Time, error) {
	// check if value = (other)
	if value == other {
		return nil, fmt.Errorf("field %q has time value %q, skip", field, other)
	}

	layout, ok := timeFormats[field]
	if !ok {
		return nil, fmt.Errorf("field %q time format isn't supported", field)
	}

	val, err := time.Parse(layout, value)
	if err != nil {
		return nil, err
	}

	return &val, nil
}
