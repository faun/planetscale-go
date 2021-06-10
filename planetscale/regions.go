package planetscale

import (
	"context"
	"net/http"

	"github.com/pkg/errors"
)

const regionsAPIPath = "v1/regions"

type Region struct {
	Slug string `json:"slug"`
	Name string `json:"display_name"`
}

type regionsResponse struct {
	Regions []*Region `json:"data"`
}

type RegionsService interface {
	List(ctx context.Context) ([]*Region, error)
}

type regionsService struct {
	client *Client
}

var _ RegionsService = &regionsService{}

func NewRegionsSevice(client *Client) *regionsService {
	return &regionsService{
		client: client,
	}
}

func (r *regionsService) List(ctx context.Context) ([]*Region, error) {
	req, err := r.client.newRequest(http.MethodGet, regionsAPIPath, nil)
	if err != nil {
		return nil, errors.Wrap(err, "error creating request for list regions")
	}

	regionsResponse := &regionsResponse{}
	if err := r.client.do(ctx, req, &regionsResponse); err != nil {
		return nil, err
	}

	return regionsResponse.Regions, nil
}
