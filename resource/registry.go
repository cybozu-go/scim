package resource

import (
	"reflect"
	"sync"
)

// The registry contains the mapping from schema URI to a Go object
type Registry struct {
	mu      sync.RWMutex
	objects map[string]reflect.Type
}

func (r *Registry) Register(uri string, data interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.objects[uri] = reflect.TypeOf(data)
}

func (r *Registry) Get(uri string) (reflect.Type, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rt, ok := r.objects[uri]
	return rt, ok
}

var registry = &Registry{
	objects: make(map[string]reflect.Type),
}

func RegisterExtension(uri string, data interface{}) {
	registry.Register(uri, data)
}
