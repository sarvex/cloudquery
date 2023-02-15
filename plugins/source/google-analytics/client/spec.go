package client

import (
	"fmt"
	"time"
)

type Spec struct {
	APIKey    string `json:"api_key,omitempty"`
	ViewID    string `json:"view_id,omitempty"`
	StartDate string `json:"start_date,omitempty"`
}

const layout = "2006-01-02"

func (s *Spec) setDefaults() {
	if len(s.StartDate) == 0 {
		// date 30 days prior
		s.StartDate = time.Now().UTC().Add(-30 * 24 * time.Hour).Format(layout)
	}
}

func (s *Spec) validate() error {
	if len(s.ViewID) == 0 {
		return fmt.Errorf(`required field "view_id" is missing`)
	}

	_, err := time.Parse(layout, s.StartDate)
	if err != nil {
		return fmt.Errorf(`"start_date" has to be in %q format, got %q: %w`, layout, s.StartDate, err)
	}

	return nil
}
