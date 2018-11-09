package endpoint

import (
	"context"

	"github.com/graniticio/granitic/ws"
	"github.com/graniticio/granitic/logging"
)

type StatusLogic struct {
	Log logging.Logger
}

func (st *StatusLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
	for name, headers := range req.UnderlyingHTTP.Request.Header {
		for _, h := range headers {
			st.Log.LogInfof("[Header] %v: %v", name, h)
		}
	}

	res.Body = "All Good!"
}
