package godc

//ArticleSearchData contains post info including ser_pos URL
type ArticleSearchData struct {
	Articles []ArticleData `json:"articles"`
	NextPos  string        `json:"nextPos"`
}

//Writer contains writer information.
type Writer struct {
	Name       string `json:"name"`
	IsSignedIn bool   `json:"isSignedIn"`
	Identity   string `json:"identity"`
}

//ArticleData contains post(list) information.
type ArticleData struct {
	URL        string `json:"url"`
	Title      string `json:"title"`
	Writer     Writer `json:"writer"`
	Type       string `json:"type"`
	Tag        string `json:"tag"`
	ReplyCount string `json:"replyCount"`
	Timestamp  string `json:"timestamp"`
	ViewCount  string `json:"viewCount"`
	UpVote     string `json:"upVote"`
}

//ArticleBody contains article information.
type ArticleBody struct {
	Title       string  `json:"title"`
	Writer      Writer  `json:"writer"`
	Timestamp   string  `json:"timestamp"`
	ViewCounter string  `json:"viewCount"`
	ReplyCount  string  `json:"replyCount"`
	Body        string  `json:"body"`
	UpVote      string  `json:"upVote"`
	DownVote    string  `json:"downVote"`
	Replies     []Reply `json:"replies"`
}

//Reply contains reply data of articleBody.
type Reply struct {
	Writer    Writer `json:"writer"`
	Type      string `json:"type"`
	Body      string `json:"body"`
	Timestamp string `json:"timestamp"`
}

//GallInfo contains gallery info used by FetchMajor/MinorGallList.
type GallInfo struct {
	Category   string `json:"category"`
	Name       string `json:"name"`
	KoName     string `json:"ko_name"`
	Manager    string `json:"manager"`
	SubManager string `json:"submanager"`
	No         string `json:"no"`
}
