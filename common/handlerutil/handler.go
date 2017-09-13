package handlerutil

import (
	"encoding/json"
	"net/http"

	"golang.org/x/net/context"
)

const (
	//Default request body size in MB.
	MaxRequestSize = 32 * 1024 * 1024 //32MB.
)

var DecodeJSON = func(ctx context.Context, w http.ResponseWriter, r *http.Request,
	v interface{}) error {

	err := json.NewDecoder(r.Body).Decode(v)
	defer r.Body.Close()
	return err
}
