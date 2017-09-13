package main

import (
	"net/http"

	"github.com/rickychandra/goboil/common/config"
	"github.com/rickychandra/goboil/common/ctxutil"
	"github.com/rickychandra/goboil/common/handlerutil"
	"github.com/rickychandra/goboil/common/logger"
	"github.com/rickychandra/goboil/common/stringutil"
	"golang.org/x/net/context"
)

func main() {
	logger.Init()

	http.HandleFunc("/hello", ctxutil.WithContext(helloHandler))
	http.HandleFunc("/json", ctxutil.WithContext(testHandler))

	http.ListenAndServe(config.Get(config.MyHandlerPort), new(APIHandler))
}

type APIHandler struct{}

func (h *APIHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, handlerutil.MaxRequestSize)
	http.DefaultServeMux.ServeHTTP(w, r)
}

var helloHandler = func(_ context.Context, w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Hello world!"))
}

type jsonReq struct {
	Name       string `json:"name"`
	BodyHeight int    `json:"bodyHeight"`
}

type jsonRes struct {
	OK bool `json:"ok"`
}

var testHandler = func(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	req := new(jsonReq)
	err := handlerutil.DecodeJSON(ctx, w, r, req)
	if err != nil {
		logger.Error(ctx, err)
		return
	}

	logger.Info(ctx, stringutil.Sprintf("json req: %v", *req))

	res := new(jsonRes)
	res.OK = true
	handlerutil.SendJSONResponse(ctx, w, r, res)
}
