package endpoint

import (
	"context"

	"github.com/graniticio/granitic/ws"
)

func (st *StatusLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
	res.Body = "All Good!"
}
