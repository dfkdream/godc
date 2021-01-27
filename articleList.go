//Package godc reads list of posts from the DCInside Gallery.
//
//디시인사이드 갤러리 게시글 목록을 읽어오는 패키지
package godc

import (
	"errors"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

func attrGet(node *html.Node, attr string) string {
	if node == nil || len(node.Attr) < 1 {
		return ""
	}
	for _, att := range node.Attr {
		if att.Key == attr {
			return att.Val
		}
	}
	return ""
}

//FetchArticleList reads specific page of post list.
//
//게시글 목록의 지정된 페이지를 읽어옵니다
func FetchArticleList(gallID string, page int, recommend bool) ([]ArticleData, error) {
	dcpg := fetchRawArticleList(gallID, page, recommend)
	if dcpg == nil {
		return nil, errors.New("Page fetch error")
	}
	qdoc, err := goquery.NewDocumentFromReader(dcpg)
	if err != nil {
		return nil, err
	}

	adataResult := make([]ArticleData, 0)
	detailList := qdoc.Find("ul.gall-detail-lst")
	detailList.Find("li").Each(func(i int, s *goquery.Selection) {
		var article ArticleData
		article.URL, _ = s.Find("a.lt").Attr("href")
		if article.URL == "" {
			return
		}
		article.Title = s.Find("span.subjectin").Text()
		article.Type, _ = s.Find("span.sp-lst").Attr("class")
		article.ReplyCount = s.Find("span.ct").Text()
		blockInfo := s.Find("span.blockInfo")
		article.Writer.Name = blockInfo.AttrOr("data-name", "")
		article.Writer.Identity = blockInfo.AttrOr("data-info", "")
		ginfo := s.Find("ul.ginfo")
		liCnt := len(ginfo.Find("li").Nodes)
		article.Writer.IsSignedIn = ginfo.Find("span").HasClass("sp-nick")
		ginfo.Find("li").Each(func(i int, s *goquery.Selection) {
			if liCnt == 5 {
				switch i {
				case 0:
					article.Tag = s.Text()
				case 2:
					article.Timestamp = s.Text()
				case 3:
					fmt.Sscanf(s.Text(), "조회 %s", &article.ViewCount)
				case 4:
					fmt.Sscanf(s.Text(), "추천 %s", &article.UpVote)
				}
			} else {
				switch i {
				case 1:
					article.Timestamp = s.Text()
				case 2:
					fmt.Sscanf(s.Text(), "조회 %s", &article.ViewCount)
				case 3:
					fmt.Sscanf(s.Text(), "추천 %s", &article.UpVote)
				}
			}
		})
		adataResult = append(adataResult, article)
	})
	return adataResult, nil
}
