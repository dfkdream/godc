package godc

import "testing"

func TestFetchArticleSearch(t *testing.T) {
	data, err := FetchArticleSearch("github", "1", "오픈소스", "all", "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	for i, d := range data.Articles {
		if !isArticleDataOK(d) {
			t.Errorf("required field is missing: %d %+v", i, d)
		}
	}

	if data.NextPos == "" {
		t.Errorf("data.NextPos is missing")
	}
}
