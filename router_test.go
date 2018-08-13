package main

import (
	"net/http"
	"testing"
)

func TestNewRoute(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {}
	r := NewRoute(http.MethodPost, "/users", f)

	e := &Route{http.MethodPost, "/users", f}

	if r.Path != e.Path || r.Method != e.Method {
		t.Errorf("unexoect route: actual=[%+v] expect=[%+v]", *r, *e)
	}
}

func TestNewRouter(t *testing.T) {
	if NewRouter() != nil {
		t.Error("unexoect router is nil")
	}
}
