package plugin

import (
	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/client"
	"github.com/cloudquery/plugin-sdk/plugins/source"
)

var Version = "Development"

func GoogleAnalytics() *source.Plugin {
	return source.NewPlugin(
		"google-analytics",
		Version,
		tables(),
		client.Configure,
	)
}
