package godc

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

//https://stackoverflow.com/questions/30109061/golang-parse-html-extract-all-content-with-body-body-tags
func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func removeScript(doc string) string {
	vaildScript := regexp.MustCompile(`<script\b[^>]*>([\s\S]*?)</script>`)
	return vaildScript.ReplaceAllString(doc, "")
}

//FetchArticleData reads article data of specific URL
//
//지정한 URL의 게시물을 읽어옵니다.
func FetchArticleData(URL string) (*ArticleBody, error) {
	dcArticle := fetchURL(URL)
	if dcArticle == nil {
		return nil, errors.New("Page fetch error")
	}
	qdoc, err := goquery.NewDocumentFromReader(dcArticle)
	if err != nil {
		return nil, err
	}

	result := ArticleBody{}

	result.IsNew = "true" //구버전 구분용

	//gallview-tit-box 처리파트 시작
	headerdiv := qdoc.Find("div.gallview-tit-box")
	result.Title = strings.TrimSpace(headerdiv.Find("span.tit").Text())
	infoul := headerdiv.Find("ul.ginfo2")
	infoul.Find("li").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			result.Name = s.Text()
		case 1:
			result.Timestamp = s.Text()
		}
	})
	result.GallogURL, _ = headerdiv.Find("a.btn-line-gray").Attr("href")
	//gallview-tit-box 처리파트 종료

	//gallview-thum-btm-inner 처리파트 시작
	articleInner := qdoc.Find("div.gall-thum-btm-inner")
	//조회수-추천-댓글수 처리파트 시작
	aiGinfo2 := articleInner.Find("ul.ginfo2")
	aiGinfo2.Find("li").Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0:
			fmt.Sscanf(s.Text(), "조회수 %s", &result.ViewCounter)
		case 1:
			fmt.Sscanf(s.Text(), "추천 %s", &result.UpVote)
		case 2:
			result.ReplyCount = s.Find("span.point-red").Text()
		}
	})
	//조회수-추천-댓글수 처리파트 종료

	//thum-txt(본문) 처리파트 시작
	rawThumtxt, _ := articleInner.Find("div.thum-txtin").Html()
	result.Body = removeScript(strings.TrimSpace(rawThumtxt))
	//thum-txt(본문) 처리파트 종료

	//추천-비추천 처리파트 시작
	result.UpVote += "/" + articleInner.Find("span#recomm_btn_member.num").Text() //고닉추
	result.DownVote = articleInner.Find("span#nonrecomm_btn.no-ct").Text()        //비추
	//추천-비추천 처리파트 종료
	//gallview-thum-btm-inner 처리파트 종료

	//all-comment-list(댓글) 처리파트 시작
	comments := make([]Reply, 0)

	allcomment := qdoc.Find("ul.all-comment-lst")
	allcomment.Find("li").Each(func(i int, s *goquery.Selection) {
		if _, exist := s.Attr("id"); !exist {
			return
		}
		var comment Reply
		if v, _ := s.Attr("class"); v == "comment" {
			comment.Type = "reply"
		} else if v == "comment-add " {
			comment.Type = "re-reply"
		}
		comment.URL, _ = s.Find("a.nick").Attr("href")
		s.Find("a.nick").Contents().Each(func(j int, t *goquery.Selection) {
			if goquery.NodeName(t) == "#text" {
				comment.Name = t.Text()
			}
		})
		comment.ID = s.Find("span.blockCommentId").Text()
		comment.IP = s.Find("span.ip").Text()
		comment.Timestamp = s.Find("span.date").Text()
		commentbHTML, _ := s.Find("p.txt").Html()
		comment.Body = removeScript(strings.TrimSpace(commentbHTML))
		comments = append(comments, comment)
	})
	//all-comment-list(댓글) 처리파트 종료
	result.Replies = comments

	return &result, nil

}
