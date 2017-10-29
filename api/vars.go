package main

import (
	"net/http"
	"sync"
)

var vars map[*http.Request]map[string]interface{}
var varsLock sync.RWMutex

// OpenVars holds variable for particular request.
func OpenVars(r *http.Request) {
	varsLock.Lock()
	if vars == nil {
		vars = map[*http.Request]map[string]interface{}{}
	}
	vars[r] = map[string]interface{}{}
	varsLock.Unlock()
}

// CloseVars cleans up memory used by map.
func CloseVars(r *http.Request) {
	varsLock.Lock()
	delete(vars, r)
	varsLock.Unlock()
}

// GetVar gets values for a key in map.
func GetVar(r *http.Request, key string) interface{} {
	varsLock.RLock()
	value := vars[r][key]
	varsLock.RLock()
	return value
}

// SetVar sets variable in map with key & val.
func SetVar(r *http.Request, key string, value interface{}) {
	varsLock.Lock()
	vars[r][key] = value
	varsLock.Unlock()
}
