package endpoint

import (
  "context"
  "github.com/graniticio/granitic/ws"
  "github.com/satori/go.uuid"
)

type WorkerLogic struct {
}

type WorkerCreateLogic struct {
}

type WorkerRequest struct {
  Id string
}

type WorkerCreateRequest struct {
  Id uuid.UUID
  FirstName string 
  LastName string
  Email string
  Address string
}

func (al *WorkerLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

  wr := req.RequestBody.(*WorkerRequest)

  res.Body = wr.Id
}

func (al *WorkerCreateLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {

  wr := req.RequestBody.(*WorkerCreateRequest)

  wr.Id = generateUid()
  res.Body = wr
}

func (al *WorkerLogic) UnmarshallTarget() interface{} {
  return new(WorkerRequest)
}


func (al *WorkerCreateLogic) UnmarshallTarget() interface{} {
  return new(WorkerCreateRequest)
}

func generateUid() uuid.UUID {
  uid := uuid.Must(uuid.NewV4())
  return uid
}

type WorkerDetail struct {
  Name string
}

