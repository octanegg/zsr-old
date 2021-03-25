package ballchasing

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff"
)

const (
	// BaseURL .
	BaseURL = "https://ballchasing.com/api"
)

type client struct {
	AuthToken string
}

// Client .
type Client interface {
	GetReplay(string) (*Replay, error)
	ListReplays(map[string][]string) (*Replays, error)
}

// New .
func New(authToken string) Client {
	return &client{authToken}
}

func (b *client) GetReplay(id string) (*Replay, error) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, BaseURL+"/replays/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", b.AuthToken)

	var replay Replay
	operation := func() error {
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusTooManyRequests {
			return errors.New(http.StatusText(http.StatusTooManyRequests))
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(body, &replay); err != nil {
			return err
		}

		return nil
	}

	if err := backoff.Retry(operation, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	return &replay, nil
}

func (b *client) ListReplays(params map[string][]string) (*Replays, error) {
	req, err := http.NewRequestWithContext(context.TODO(), http.MethodGet, BaseURL+"/replays", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", b.AuthToken)
	req.URL.RawQuery = url.Values(params).Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var replays Replays
	if err := json.Unmarshal(body, &replays); err != nil {
		return nil, err
	}

	return &replays, nil
}
