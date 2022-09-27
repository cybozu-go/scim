package resource

import (
	"reflect"
	"sync"
)

// The registry contains the mapping from schema URI to a Go object
type Registry struct {
	mu      sync.RWMutex
	urimap  map[string]reflect.Type
	namemap map[string]reflect.Type
}

func (r *Registry) Register(name, uri string, data interface{}) {
	r.mu.Lock()
	defer r.mu.Unlock()
	rt := reflect.TypeOf(data)
	r.namemap[name] = rt
	if uri != "" {
		r.urimap[uri] = rt
	}
}

func (r *Registry) LookupByURI(uri string) (interface{}, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rt, ok := r.urimap[uri]
	if !ok {
		return nil, false
	}
	return reflect.New(rt).Interface(), true
}

func (r *Registry) LookupByName(name string) (interface{}, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	rt, ok := r.namemap[name]
	if !ok {
		return nil, false
	}
	return reflect.New(rt).Interface(), true
}

var registry = &Registry{
	urimap:  make(map[string]reflect.Type),
	namemap: make(map[string]reflect.Type),
}

func Register(name, uri string, data interface{}) {
	registry.Register(name, uri, data)
}
