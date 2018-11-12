package endpoint

import (
  "context"
  "fmt"
  "github.com/graniticio/granitic/ws"
  "github.com/satori/go.uuid"
  "worker-management/dbms"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/graniticio/granitic/logging"
  "net/http"
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

type WorkerCreateRequest map[string]interface{}

type DynamoWorkerCreateRequest struct {
  Id string
  Document interface{}
  Version int
  Schema int
}

func (wl *WorkerLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
  dynamoClient := wl.DBManager.Client()

  wr := req.RequestBody.(*WorkerRequest)

  key, err := dynamodbattribute.MarshalMap(wr)
  wl.Log.LogInfof("%v", key)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    res.HttpStatus = http.StatusInternalServerError
    return
  }
  result, err := dynamoClient.GetWorker(key)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    res.HttpStatus = http.StatusInternalServerError
    return
  }

  var worker map[string]interface{}

  err = dynamodbattribute.UnmarshalMap(result.Item, &worker)

  if err != nil {
    wl.Log.LogErrorf("%v", err)
    res.HttpStatus = http.StatusInternalServerError
    return
    } else if _, ok := worker["Id"].(string); !ok {
    res.HttpStatus = http.StatusNotFound
    return
  } else if _, ok := worker["Document"].(map[string]interface{}); !ok {
    worker["Document"] = make(map[string]interface{})
  }

  responseBody := worker["Document"].(map[string]interface{})
  responseBody["Id"] = worker["Id"].(string)
  res.Body = responseBody
}

func (wl *WorkerCreateLogic) Process(ctx context.Context, req *ws.WsRequest, res *ws.WsResponse) {
  dynamoClient := wl.DBManager.Client()
  wr := req.RequestBody.(*WorkerCreateRequest)
  dwr := DynamoWorkerCreateRequest{Document: wr, Version: 1, Schema: 1}
  dwr.Id = generateUid()
  
  item, err := dynamodbattribute.MarshalMap(dwr)

  if err != nil {
    wl.Log.LogErrorf("%v", err)
    res.HttpStatus = http.StatusInternalServerError
    return
  }
  wl.Log.LogInfof("%v", item)
  err = dynamoClient.CreateWorker(item)
  if err != nil {
    wl.Log.LogErrorf("%v", err)
    res.HttpStatus = http.StatusInternalServerError
    return
  }

  res.Body = WorkerRequest{Id: dwr.Id}
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
