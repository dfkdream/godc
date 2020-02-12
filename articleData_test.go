package godc

import "testing"

func TestFetchArticleData(t *testing.T) {
	body, err := FetchArticleData("https://m.dcinside.com/board/github/6062?page=1")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Logf("%+v\n", body)
}
