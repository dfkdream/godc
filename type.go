package godc

//ArticleSearchData contains post info including ser_pos URL
type ArticleSearchData struct {
	Articles []ArticleData
	NextPos  string
}

//ArticleData contains post(list) informations.
type ArticleData struct {
	URL         string
	Title       string
	Type        string
	ReplyCount  string
	Name        string
	Timestamp   string
	ViewCounter string
	UpVote      string
	WriterID    string
}

//ArticleBody contains article informations.
type ArticleBody struct {
	IsNew       string
	Title       string
	Name        string
	IP          string
	GallogURL   string
	Timestamp   string
	ViewCounter string
	ReplyCount  string
	Body        string
	UpVote      string
	DownVote    string
	Replies     []Reply
}

//Reply contains reply data of articleBody.
type Reply struct {
	URL       string
	Name      string
	ID        string
	IP        string
	Type      string
	Body      string
	Timestamp string
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
