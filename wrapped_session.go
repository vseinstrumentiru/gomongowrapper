package gomongowrapper

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Session struct {
	mongo.Session
}

var _ mongo.Session = (*Session)(nil)

func (ws *Session) EndSession(ctx context.Context) {
	ws.Session.EndSession(ctx)
}

func (ws *Session) StartTransaction(topts ...*options.TransactionOptions) error {
	return handle(ws.Session.StartTransaction(topts...))
}

func (ws *Session) AbortTransaction(ctx context.Context) error {
	return handle(ws.Session.AbortTransaction(ctx))
}

func (ws *Session) CommitTransaction(ctx context.Context) error {
	return handle(ws.Session.CommitTransaction(ctx))
}

func (ws *Session) ClusterTime() bson.Raw {
	return ws.Session.ClusterTime()
}

func (ws *Session) AdvanceClusterTime(br bson.Raw) error {
	return handle(ws.Session.AdvanceClusterTime(br))
}

func (ws *Session) OperationTime() *primitive.Timestamp {
	return ws.Session.OperationTime()
}

func (ws *Session) AdvanceOperationTime(pt *primitive.Timestamp) error {
	return handle(ws.Session.AdvanceOperationTime(pt))
}
