package sotw

import (
	"context"
	"reflect"

	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v3"
	"github.com/envoyproxy/go-control-plane/pkg/cache/types"
	"github.com/envoyproxy/go-control-plane/pkg/cache/v3"
)

// watches for all xDS resource types
type watches struct {
	responders map[string]*watch

	// indexes is a list of indexes for each dynamic select case which match to a watch
	cases []reflect.SelectCase
}

// newWatches creates and initializes watches.
func newWatches() watches {
	return watches{
		responders: make(map[string]*watch, int(types.UnknownType)),
		cases:      make([]reflect.SelectCase, 0),
	}
}

// addWatch creates a new watch entry in the watches map.
// Watches are sorted by typeURL.
func (w *watches) addWatch(typeURL string, watch *watch) {
	w.responders[typeURL] = watch
}

// close all open watches
func (w *watches) close() {
	for _, watch := range w.responders {
		watch.close()
	}
}

// recomputeWatches rebuilds the known list of dynamic channels if needed
func (w *watches) recompute(ctx context.Context, req <-chan *discovery.DiscoveryRequest) {
	w.cases = w.cases[0:]
	w.cases = append(w.cases,
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(ctx.Done()),
		}, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(req),
		},
	)

	for _, watch := range w.responders {
		w.cases = append(w.cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(watch.response),
		})
	}
}

// watch contains the necessary modifiables for receiving resource responses
type watch struct {
	cancel   func()
	nonce    string
	response chan cache.Response
}

// close cancels an open watch
func (w *watch) close() {
	if w.cancel != nil {
		w.cancel()
	}
}