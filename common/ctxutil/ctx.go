package ctxutil

import (
	"net/http"

	"github.com/rickychandra/goboil/common/stringutil"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

const (
	ctxRequestId string = "ctx-request-id"
)

//HandlerWithCtxFunc should be implemented by all API handlers.
type HandlerWithCtxFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request)

//ExtractRequestID extracts request id from context. May return empty string.
var ExtractRequestID = func(ctx context.Context) string {
	return GetMetadata(ctx, ctxRequestId)
}

//GetMetadata gets value from incoming and outgoing context with key. If key is present in both
//places, return the one in incoming context. May return empty string.
var GetMetadata = func(ctx context.Context, key string) string {
	if ctx == nil {
		return ""
	}

	mdIn, okIn := metadata.FromIncomingContext(ctx)
	mdOut, okOut := metadata.FromOutgoingContext(ctx)

	var res string
	if okOut {
		vs := mdOut[key]
		if len(vs) > 0 {
			res = vs[0]
		} else {
			res = ""
		}
	}
	if okIn {
		vs := mdIn[key]
		if len(vs) > 0 {
			res = vs[0]
		} else {
			res = ""
		}
	}

	return res
}

//PreprocessAPICtx copies request id from incoming context to outgoing context.
var PreprocessAPICtx = func(ctx context.Context) context.Context {
	reqID := ExtractRequestID(ctx)
	outMD := metadata.Pairs(ctxRequestId, reqID)
	return metadata.NewOutgoingContext(ctx, outMD)
}

//WithContext adds context to http.HandlerFunc.
var WithContext = func(f HandlerWithCtxFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		md := metadata.Pairs(ctxRequestId, stringutil.Random(25, true, true))
		ctx = metadata.NewOutgoingContext(ctx, md)
		f(ctx, w, r)
	}
}
