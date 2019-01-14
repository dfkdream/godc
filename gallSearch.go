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
		dcpg = fetchURL(fmt.Sprintf("http://m.dcinside.com/board/%s?serval=%s&s_type=%s&page=%s", gallID, query, searchType, page))
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
		URL, _ := s.Find("a.lt").Attr("href")
		if URL == "" {
			return
		}
		Title, _ := s.Find("span.detail-txt").Html()
		Type, _ := s.Find("span.sp-lst").Attr("class")
		ReplyCount := s.Find("span.ct").Text()
		Tag := ""
		Name := ""
		Timestamp := ""
		ViewCounter := ""
		UpVote := ""
		WriterID := s.Find("span.blockInfo").Text()
		ginfo := s.Find("ul.ginfo")
		liCnt := len(ginfo.Find("li").Nodes)
		ginfo.Find("li").Each(func(i int, s *goquery.Selection) {
			if liCnt == 5 {
				switch i {
				case 0:
					Tag = s.Text()
				case 1:
					Name = s.Text()
				case 2:
					Timestamp = s.Text()
				case 3:
					fmt.Sscanf(s.Text(), "조회 %s", &ViewCounter)
				case 4:
					fmt.Sscanf(s.Text(), "추천 %s", &UpVote)
				}
			} else {
				switch i {
				case 0:
					Name, _ = s.Html()
				case 1:
					Timestamp = s.Text()
				case 2:
					fmt.Sscanf(s.Text(), "조회 %s", &ViewCounter)
				case 3:
					fmt.Sscanf(s.Text(), "추천 %s", &UpVote)
				}
			}
		})
		asdataResult = append(asdataResult, ArticleData{URL, Title, Type, Tag, ReplyCount, Name, Timestamp, ViewCounter, UpVote, WriterID})
	})

	nextURL, _ := qdoc.Find("a.next").Attr("href")

	return &ArticleSearchData{Articles: asdataResult, NextPos: nextURL}, nil
}
