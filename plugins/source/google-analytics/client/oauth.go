package client

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net"
	"net/http"
	"os/exec"

	"github.com/rs/zerolog"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/analyticsreporting/v4"
)

func getTokenSource(ctx context.Context, l *zerolog.Logger) (oauth2.TokenSource, error) {
	lst, err := net.Listen("tcp4", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	l.Info().Str("addr", lst.Addr().String()).Msg("will listen")

	config := &oauth2.Config{
		ClientID:     "cloudquery-google-analytics-source-plugin",
		ClientSecret: "we_really_dont_care",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://" + lst.Addr().String(),
		Scopes:       []string{analyticsreporting.AnalyticsReadonlyScope},
	}

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)

	handler := &oauthHandler{
		state:  state,
		err:    make(chan error),
		logger: l,
	}

	srv := http.Server{Handler: handler}

	go func() {
		defer srv.Close()
		exec.CommandContext(ctx, "open", config.AuthCodeURL(state, oauth2.AccessTypeOffline)).Run()
		err = <-handler.err
	}()

	if serveErr := srv.Serve(lst); serveErr != http.ErrServerClosed {
		return nil, serveErr
	}

	l.Info().Err(err).Msg("served")
	if err != nil {
		return nil, err
	}

	// we have exchange token now
	token, err := config.Exchange(ctx, handler.code, oauth2.AccessTypeOffline)
	l.Info().Err(err).Str("token", token.AccessToken).Msg("got tok")
	if err != nil {
		return nil, err
	}

	return config.TokenSource(context.Background(), token), nil
}

type oauthHandler struct {
	state  string
	err    chan error
	code   string
	logger *zerolog.Logger
}

func (o *oauthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.logger.Info().Str("code", r.FormValue("code")).Msg("got code")
	//TODO implement me
	panic("implement me")
}
