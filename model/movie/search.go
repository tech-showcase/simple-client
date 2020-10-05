package movie

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/tech-showcase/entertainment-service/helper"
	"net/http"
	"net/url"
	"strconv"
)

type (
	SearchMovieRequest struct {
		Keyword    string `json:"keyword"`
		PageNumber int    `json:"page_number"`
	}
	SearchMovieResponse struct {
		ListPerPage
	}
)

func makeSearchMovieClientEndpoint(movieServiceURL *url.URL) endpoint.Endpoint {
	searchMovieURL, _ := helper.JoinURL(movieServiceURL, "/movie")

	searchMovieClientEndpoint := httptransport.NewClient(
		http.MethodPost,
		searchMovieURL,
		encodeSearchMovieHTTPRequest,
		decodeSearchMovieHTTPResponse,
	).Endpoint()

	return searchMovieClientEndpoint
}

func encodeSearchMovieHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	if req, ok := request.(SearchMovieRequest); ok {
		q := r.URL.Query()
		q.Add("keyword", req.Keyword)
		q.Add("page_number", strconv.Itoa(req.PageNumber))
		r.URL.RawQuery = q.Encode()

		return nil
	} else {
		return errors.New("request format is wrong")
	}
}

func decodeSearchMovieHTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var req SearchMovieResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
