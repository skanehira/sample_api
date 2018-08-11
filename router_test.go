package main

import (
	"reflect"
	"testing"
)

func TestGetQueryParams(t *testing.T) {

	validPath := []struct {
		path  string
		query QueryParams
	}{
		{"/users?id=neko?pass=pass", QueryParams{"id": "neko", "pass": "pass"}},
		{"/users?id=neko", QueryParams{"id": "neko"}},
		{"/users?id=?pass=", QueryParams{"id": "", "pass": ""}},
		{"/users?id?pass", QueryParams{"id": "", "pass": ""}},
	}

	// invalidPath := []map[string]QueryParams{
	// 	{"/users?id=?pass": QueryParams{"id": "", "pass": ""}},
	// 	{"/users?id?pass=": QueryParams{"id": "", "pass": ""}},
	// 	{"/users??pass=": QueryParams{"pass": ""}},
	// 	{"/userspass=": QueryParams{"": ""}},
	// }

	for _, except := range validPath {
		t.Logf("test data:%s", except.path)
		query := getQueryParams(except.path)
		if !reflect.DeepEqual(except.query, query) {
			t.Errorf("unexpect params: actual=[%+v] expect=[%+v]", query, except.query)
		}
	}
}
