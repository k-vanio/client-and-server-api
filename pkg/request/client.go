package request

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/x-vanio/client-and-server-api/pkg/dto"
)

type HTTPClient interface {
	ByDollarQuote(url string) (*dto.Quote, error)
}

type client struct {
	request interface {
		Do(req *http.Request) (*http.Response, error)
	}
	timeout time.Duration
}

func (c *client) ByDollarQuote(url string) (*dto.Quote, error) {
	// request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// context
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	// timeout
	req = req.WithContext(ctx)

	resp, err := c.request.Do(req)
	if err != nil {
		if ctxErr := ctx.Err(); ctxErr == context.DeadlineExceeded {
			log.Println(ctxErr.Error())
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := &dto.Quote{}
	if err = json.Unmarshal(body, response); err != nil {
		return nil, err
	}

	return response, nil

}

func NewClient(
	request interface {
		Do(req *http.Request) (*http.Response, error)
	},
	timeout time.Duration,
) *client {
	return &client{request: request, timeout: timeout}
}
