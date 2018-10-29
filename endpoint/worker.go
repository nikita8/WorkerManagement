package endpoint

import (
  "context"
  "github.com/graniticio/granitic/ws"
  "github.com/satori/go.uuid"
  "worker-management/dbms"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/graniticio/granitic/logging"
  "github.com/graniticio/granitic/types"
)

type WorkerLogic struct {
  DBManager *dbms.ClientManager
  Log logging.Logger
}

type WorkerCreateLogic struct {
  DBManager *dbms.ClientManager
  Log logging.Logger
}

type WorkerRequest struct {
  Id string
}

type WorkerCreateRequest struct {
  Id string
  FirstName *types.NilableString 
  LastName *types.NilableString
  Email *types.NilableString
  Address *types.NilableString
}

func (wl *WorkerLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
  dynamoClient := wl.DBManager.Client()

  wr := req.RequestBody.(*WorkerRequest)

  key, err := dynamodbattribute.MarshalMap(wr)

  if err != nil {
    wl.Log.LogErrorf("%v", err)
    return
  }
  result, err := dynamoClient.GetWorker(key)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    return
  }

  worker := WorkerCreateRequest{}

  err = dynamodbattribute.UnmarshalMap(result.Item, &worker)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    return
  }

  res.Body = worker
}

func (wl *WorkerCreateLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
  dynamoClient := wl.DBManager.Client()
  wr := req.RequestBody.(*WorkerCreateRequest)
  wr.Id = generateUid()

  item, err := dynamodbattribute.MarshalMap(wr)

  if err != nil {
    wl.Log.LogErrorf("%v", err)
    return
  }

  err = dynamoClient.CreateWorker(item)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    return
  }

  res.Body = wr.Id
}

func (wl *WorkerLogic) UnmarshallTarget() interface{} {
  return new(WorkerRequest)
}


func (wl *WorkerCreateLogic) UnmarshallTarget() interface{} {
  return new(WorkerCreateRequest)
}

func generateUid() string {
  uid := uuid.Must(uuid.NewV4()).String()
  return uid
}

type WorkerDetail struct {
  Name string
}

