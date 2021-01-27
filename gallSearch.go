package godc

import (
	"errors"
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

//FetchArticleSearch reads specific page of post search result.
//
//페이지 검색의 지정한 페이지를 읽어옵니다.
func FetchArticleSearch(gallID string, page string, query string, searchType string, next string) (*ArticleSearchData, error) {
	var dcpg io.ReadCloser
	if next == "" || next == "0" {
		dcpg = fetchURL(fmt.Sprintf("https://m.dcinside.com/board/%s?serval=%s&s_type=%s&page=%s", gallID, query, searchType, page))
	} else {
		dcpg = fetchURL(next)
	}
	if dcpg == nil {
		return nil, errors.New("Page fetch error")
	}
	qdoc, err := goquery.NewDocumentFromReader(dcpg)
	if err != nil {
		return nil, err
	}

	asdataResult := make([]ArticleData, 0)

	detailList := qdoc.Find("ul.gall-detail-lst")
	detailList.Find("li").Each(func(i int, s *goquery.Selection) {
		var article ArticleData
		article.URL, _ = s.Find("a.lt").Attr("href")
		if article.URL == "" {
			return
		}
		article.Title, _ = s.Find("span.subjectin").Html()
		article.Type, _ = s.Find("span.sp-lst").Attr("class")
		article.ReplyCount = s.Find("span.ct").Text()
		blockInfo := s.Find("span.blockInfo")
		article.Writer.Name = blockInfo.AttrOr("data-name", "")
		article.Writer.Identity = blockInfo.AttrOr("data-info", "")
		ginfo := s.Find("ul.ginfo")
		liCnt := len(ginfo.Find("li").Nodes)
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
		asdataResult = append(asdataResult, article)
	})

	nextURL, _ := qdoc.Find("a.next").Attr("href")

	return &ArticleSearchData{Articles: asdataResult, NextPos: nextURL}, nil
}
