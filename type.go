package godc

//ArticleData contains post URL, post title, username.
//
//게시글 주소(URL), 게시글 제목(Title), 닉네임(고닉/유동 구분 안됨)(Name)
type ArticleData struct {
	URL   string
	Title string
	Name  string
}
