package main

import (
	"context"

	"github.com/MagicBowen/microservice/examples/services/utils/tracing"
)

type tracingDB struct {
	tracer *tracing.ServiceTracer
	db entityDB
}

func newTracingDB(tracer *tracing.ServiceTracer, db entityDB) *tracingDB {
	return &tracingDB{tracer : tracer, db : db}
}

func (tDB *tracingDB) traceMongoDB(
	ctx context.Context, 
	action string,
	) (context.Context, spanFinish) {
	return traceProcess(ctx, tDB.tracer, "mongoDB " + action, "", "")
}


func (tDB *tracingDB) findOne(ctx context.Context, filter interface{}, value interface{}) error{
	_, spanFinish := tDB.traceMongoDB(ctx, "findOne")
	defer spanFinish()
	return tDB.db.findOne(ctx, filter, value)
}

func (tDB *tracingDB)insertOne(ctx context.Context, value interface{}) error{
	_, spanFinish := tDB.traceMongoDB(ctx, "insertOne")
	defer spanFinish()
	return tDB.db.insertOne(ctx, value)
}

func (tDB *tracingDB)updateOne(ctx context.Context, filter interface{}, update interface{}) error{
	_, spanFinish := tDB.traceMongoDB(ctx, "updateOne")
	defer spanFinish()
	return tDB.db.updateOne(ctx, filter, update)
}

func (tDB *tracingDB)delete(ctx context.Context, filter interface{}) error{
	_, spanFinish := tDB.traceMongoDB(ctx, "delete")
	defer spanFinish()
	return tDB.db.delete(ctx, filter)
}

