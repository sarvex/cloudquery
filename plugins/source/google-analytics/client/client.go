package client

import (
	"context"

	"github.com/cloudquery/plugin-sdk/backend"
	"github.com/cloudquery/plugin-sdk/plugins/source"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
	"google.golang.org/api/analyticsreporting/v4"
	"google.golang.org/api/option"
)

type Client struct {
	*analyticsreporting.Service
	backend.Backend

	ViewID    string
	StartDate string

	logger zerolog.Logger
}

var _ schema.ClientMeta = (*Client)(nil)

func (c *Client) Logger() *zerolog.Logger {
	return &c.logger
}

func (c *Client) ID() string {
	return "google-analytics:view:{" + c.ViewID + "}"
}

func Configure(ctx context.Context, logger zerolog.Logger, srcSpec specs.Source, options source.Options) (schema.ClientMeta, error) {
	spec := new(Spec)
	if err := srcSpec.UnmarshalSpec(&spec); err != nil {
		return nil, err
	}

	spec.setDefaults()
	if err := spec.validate(); err != nil {
		return nil, err
	}

	tokenSource, err := spec.OAuth.getTokenSource(ctx)
	if err != nil {
		return nil, err
	}
	opts := []option.ClientOption{
		option.WithTokenSource(tokenSource),
		option.WithScopes(analyticsreporting.AnalyticsReadonlyScope),
		option.WithRequestReason("cloudquery resource fetch"),
		// we disable telemetry to boost performance and be on the same side with telemetry
		option.WithTelemetryDisabled(),
	}

	logger.Info().Msg("will create client")
	svc, err := analyticsreporting.NewService(context.Background(), opts...)
	if err != nil {
		return nil, err
	}
	logger.Info().Msg("created client")

	svc.UserAgent = "cloudquery:source-google-analytics/" + srcSpec.Version

	c := &Client{
		Service:   svc,
		Backend:   options.Backend,
		StartDate: spec.StartDate,
		ViewID:    spec.ViewID,
		logger:    logger.With().Str("plugin", "google-analytics").Str("view", spec.ViewID).Logger(),
	}

	return c, nil
}
