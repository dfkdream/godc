package godc

import (
	"testing"
)

func isArticleDataOK(d ArticleData) bool {
	if d.Title == "" ||
		d.Writer.Name == "" ||
		d.Writer.Identity == "" ||
		d.Type == "" ||
		d.Tag == "" ||
		d.Timestamp == "" {
		return false
	}
	return true
}

func TestFetchArticleList(t *testing.T) {
	for page := 1; page <= 2; page++ {
		data, err := FetchArticleList("github", page, false)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if !(len(data) > 0) {
			t.Errorf("expected len(data)>0 but got %d", len(data))
		}

		for i, d := range data {
			if !isArticleDataOK(d) {
				t.Errorf("required field is missing: %d %+v", i, d)
			}
		}
	}
}
