package main

import (
	"reflect"
	"testing"
)

func TestGetQueryParams(t *testing.T) {

	path := []struct {
		path  string
		query QueryParams
	}{
		{"/users?id=neko?pass=pass", QueryParams{"id": "neko", "pass": "pass"}},
		{"/users?id=neko", QueryParams{"id": "neko"}},
		{"/users?id=?pass=", QueryParams{"id": "", "pass": ""}},
		{"/users?id?pass", QueryParams{}},
		{"/users?id?pass=", QueryParams{"pass": ""}},
		{"/users?id=?pass", QueryParams{"id": ""}},
		{"/users??pass=", QueryParams{"pass": ""}},
		{"/userspass=", QueryParams{}},
		{"/userspass=?=?", QueryParams{}},
		{"/userspass=?=??", QueryParams{}},
		{"/userspass???", QueryParams{}},
		{"/userspass?=?=?", QueryParams{}},
	}

	for i, except := range path {
		t.Logf("test data%d : %s", i, except.path)
		query := getQueryParams(except.path)
		if !reflect.DeepEqual(except.query, query) {
			t.Errorf("unexpect params: actual=[%+v] expect=[%+v]", query, except.query)
		}
	}

}

func TestGetPath(t *testing.T) {
	path := []struct {
		path       string
		exceptPath string
	}{
		{"/users", "/users"},
		{"/", "/"},
		{"/users?id=id", "/users"},
		{"/users/:id/update", "/users"},
		{"/users/", "/users"},
	}

	for i, data := range path {
		t.Logf("test data%d : %s", i, data.path)
		p := getPath(data.path)

		if p != data.exceptPath {
			t.Errorf("unexpect path: actual=[%+v] expect=[%+v]", p, data.exceptPath)
		}
	}
}
