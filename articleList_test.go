package godc

import (
	"testing"
)

func TestFetchArticleList(t *testing.T) {
	data, err := FetchArticleList("github", 1, false)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for idx, d := range data {
		t.Log(idx, d)
	}
}
