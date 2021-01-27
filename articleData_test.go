package godc

import "testing"

func TestFetchArticleData(t *testing.T) {
	body, err := FetchArticleData("https://m.dcinside.com/board/github/18205")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	if body.Title == "" ||
		body.Writer.Name == "" || body.Writer.Identity == "" ||
		body.Timestamp == "" ||
		body.ViewCounter == "" ||
		body.ReplyCount == "" ||
		body.UpVote == "" ||
		body.DownVote == "" ||
		len(body.Replies) > 0 &&
			body.Replies[0].Timestamp == "" ||
		body.Replies[0].Writer.Name == "" ||
		body.Replies[0].Writer.Identity == "" ||
		body.Replies[0].Type == "" ||
		body.Replies[0].Body == "" {
		t.Errorf("missing required field: %+v", body)
	}
}
