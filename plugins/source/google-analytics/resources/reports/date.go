package reports

import (
	"context"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
)

const dateLayout = "2006-01-02"

func genDates(ctx context.Context, c *client.Client, table string) (<-chan time.Time, error) {
	res, err := c.Backend.Get(ctx, table, c.ID())
	if err != nil {
		return nil, err
	}

	startDate := c.StartDate
	if len(res) > 0 {
		startDate = res
	}

	// parse
	t, err := time.Parse(dateLayout, startDate)
	if err != nil {
		return nil, err
	}

	ch := make(chan time.Time)
	defer close(ch)

	today := time.Now().UTC()

	for !t.After(today) {
		ch <- t
		t.Add(24 * time.Hour)
	}

	return ch, nil
}
