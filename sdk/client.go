package sdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	baseURL string
	client  *http.Client
}

func NewClient(baseURL string, apiKEY string) *Client {

	client := http.DefaultClient
	client.Transport = NewAuthHeaderRoundTripper(apiKEY, http.DefaultTransport)

	return &Client{
		baseURL: baseURL,
		client:  http.DefaultClient,
	}
}

func (c Client) List(ctx context.Context, params *ListParams) (ListResponse, error) {
	limit := 10
	page := 1
	order := "asc"

	if params != nil {
		limit = params.Limit
		page = params.Page
		if params.DescOrder {
			order = "desc"
		}
	}

	u := fmt.Sprintf("%s/images/search?limit=%d&page=%d&order=%s", c.baseURL, limit, page, order)

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return ListResponse{}, err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return ListResponse{}, err
	}

	defer resp.Body.Close()

	if err := c.handleStatusCode(resp); err != nil {
		return ListResponse{}, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return ListResponse{}, err
	}

	var r ListResponse
	if err := json.Unmarshal(b, &r); err != nil {
		return ListResponse{}, err
	}

	return r, nil
}

func (c Client) Vote(ctx context.Context, imageID string, iLikeIt bool) error {
	u := fmt.Sprintf("%s/votes", c.baseURL)

	vb := VoteBody{
		ImageId: imageID,
	}

	if iLikeIt {
		vb.Value = 1
	} else {
		vb.Value = 0
	}

	bb, err := json.Marshal(vb)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, u, bytes.NewReader(bb))
	if err != nil {
		return err
	}

	req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}

	return c.handleStatusCode(resp)
}

func (c Client) handleStatusCode(response *http.Response) error {
	switch response.StatusCode {
	case http.StatusOK, http.StatusCreated:
		return nil
	case http.StatusNotFound:
		return ErrNotFound
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusInternalServerError:
		return ErrInternal
	}

	return ErrInternal
}
