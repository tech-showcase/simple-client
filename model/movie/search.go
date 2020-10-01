package movie

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func encodeHTTPRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

func decodeSearchMovieHTTPResponse(_ context.Context, r *http.Response) (interface{}, error) {
	var req SearchMovieResponse
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
