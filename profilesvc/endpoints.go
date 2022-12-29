package profilesvc

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostProfileEndpoint: MakePostProfileEndpoint(s),
	}
}

type Endpoints struct {
	PostProfileEndpoint endpoint.Endpoint
}

func MakePostProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postProfileRequest)
		e := s.PostProfile(ctx, req.Profile)
		return postProfileResponse{Err: e}, nil
	}
}

type postProfileRequest struct {
	Profile Profile
}

type postProfileResponse struct {
	Err error `json: "err, omitempty"`
}

func (r postProfileResponse) error() error { return r.Err }
