package main

import (
	"github.com/cloudquery/cloudquery/plugins/source/google-analytics/resources/plugin"
	"github.com/cloudquery/plugin-sdk/serve"
)

const sentryDSN = "TODO"

func main() {
	serve.Source(plugin.GoogleAnalytics(), serve.WithSourceSentryDSN(sentryDSN))
}
