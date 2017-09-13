package handlerutil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rickychandra/goboil/common/ctxutil"
	"github.com/rickychandra/goboil/common/logger"
	"golang.org/x/net/context"
)

const (
	headerRequestID = "X-Request-ID"
)

//SendBadRequestResponse writes a JSON bad request with status code Bad Request (400) to w.
var SendBadRequestResponse = func(ctx context.Context, w http.ResponseWriter, r *http.Request,
	data interface{}) {

	sendJSONResponse(ctx, http.StatusBadRequest, w, r, data)
}

//SendErrorResponse writes a JSON bad request with status code Internal Server Error (500) to w.
var SendErrorResponse = func(ctx context.Context, w http.ResponseWriter, r *http.Request,
	data interface{}) {

	sendJSONResponse(ctx, http.StatusInternalServerError, w, r, data)
}

//SendJSONResponse writes a JSON response with status code OK (200) to w.
var SendJSONResponse = func(ctx context.Context, w http.ResponseWriter, r *http.Request,
	data interface{}) {

	sendJSONResponse(ctx, http.StatusOK, w, r, data)
}

var sendJSONResponse = func(ctx context.Context, statusCode int, w http.ResponseWriter,
	r *http.Request, data interface{}) {

	logRequest(ctx, statusCode, r)

	w.Header().Set(headerRequestID, ctxutil.ExtractRequestID(ctx))
	w.Header().Set("Content-Type", "application/json")
	jsonStr, err := json.Marshal(data)
	if err != nil {
		logger.Error(ctx, err)
	}
	w.Write(jsonStr)
}

var logRequest = func(ctx context.Context, statusCode int, r *http.Request) {
	msg := fmt.Sprintf("[%d] %s", statusCode, r.URL.Path)
	logger.Info(ctx, &msg)
}
