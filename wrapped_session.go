package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type WrappedSession struct {
	mongo.Session
}

var _ mongo.Session = (*WrappedSession)(nil)

func (ws *WrappedSession) EndSession(ctx context.Context) {
	ws.Session.EndSession(ctx)
}

func (ws *WrappedSession) StartTransaction(topts ...*options.TransactionOptions) error {
	return handle(ws.Session.StartTransaction(topts...))
}

func (ws *WrappedSession) AbortTransaction(ctx context.Context) error {
	return handle(ws.Session.AbortTransaction(ctx))
}

func (ws *WrappedSession) CommitTransaction(ctx context.Context) error {
	return handle(ws.Session.CommitTransaction(ctx))
}

func (ws *WrappedSession) ClusterTime() bson.Raw {
	return ws.Session.ClusterTime()
}

func (ws *WrappedSession) AdvanceClusterTime(br bson.Raw) error {
	return handle(ws.Session.AdvanceClusterTime(br))
}

func (ws *WrappedSession) OperationTime() *primitive.Timestamp {
	return ws.Session.OperationTime()
}

func (ws *WrappedSession) AdvanceOperationTime(pt *primitive.Timestamp) error {
	return handle(ws.Session.AdvanceOperationTime(pt))
}
