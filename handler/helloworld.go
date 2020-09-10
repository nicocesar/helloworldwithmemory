package handler

import (
	"context"

	"github.com/micro/micro/v3/service/logger"
	helloworld "github.com/nicocesar/helloworldwithmemory/proto"

	gostore "github.com/micro/go-micro/v3/store"
	"github.com/micro/micro/v3/service/store"
)

// Helloworld is the most basic struct with no meanful representation
type Helloworld struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Helloworld) Call(ctx context.Context, req *helloworld.Request, rsp *helloworld.Response) error {
	logger.Info("Received Helloworld.Call request")

	var previous []*gostore.Record
	previous, _ = store.Read("previous hello")

	// I don't know why store.Read() returns an array or answers?...
	// but we'll concatenate them all
	var remaining string
	for _, p := range previous {
		remaining = remaining + string(p.Value)
	}

	rsp.Msg = "Hello5 " + req.Name + " " + remaining

	store.Write(&gostore.Record{
		Key:   "previous hello",
		Value: []byte(req.Name + " " + remaining),
	})

	return nil
}
