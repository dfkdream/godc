//Package godc reads list of posts from the DCInside Gallery.
//
//디시인사이드 갤러리 게시글 목록을 읽어오는 패키지
package godc

import (
	"errors"
	"strings"

	"golang.org/x/net/html"
)

//FetchAritcleList reads specific page of post list.
//
//게시글 목록의 지정된 페이지를 읽어옵니다
func FetchAritcleList(gallID string, page int) ([]ArticleData, error) {
	dcpg := fetchRawArticleList(gallID, page)
	if dcpg == nil {
		return nil, errors.New("Page fetch error")
	}
	doc, err := html.Parse(dcpg)
	if err != nil {
		return nil, errors.New("html parse error")
	}
	res := make([]string, 0)
	var f func(*html.Node, int)
	f = func(n *html.Node, depth int) {
		if n.FirstChild != nil && n.FirstChild.Data != "" {
			if depth == 11 || depth == 13 || depth == 14 {
				if depth == 11 {
					if n.Type == html.ElementNode && n.Data == "a" && len(n.Attr) > 0 && n.Attr[0].Key == "href" && n.Attr[0].Val != "javascript:;" {
						res = append(res, n.Attr[0].Val)
					}
				}
				if depth == 13 || depth == 14 {
					res = append(res, n.FirstChild.Data)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, depth+1)
		}
	}
	f(doc, 0)

	res = res[:len(res)-10]

	currentIndex := 0
	cuttedResult := make([][]string, 0)
	tmpResult := make([]string, 0)

	for {
		if currentIndex > len(res)-1 {
			break
		}
		tmpResult = append(tmpResult, res[currentIndex])
		if res[currentIndex] == "|" {
			tmpResult = append(tmpResult, res[currentIndex+1:currentIndex+7]...)
			currentIndex += 7
			cuttedResult = append(cuttedResult, tmpResult)
			tmpResult = make([]string, 0)
			continue
		}
		currentIndex++
	}

	processedResult := make([]ArticleData, 0)

	for _, data := range cuttedResult {
		switch len(data) {
		case 11:
			processedResult = append(processedResult,
				ArticleData{URL: data[0],
					Title:       data[1],
					ReplyCount:  "[0]",
					Name:        data[2],
					Timestamp:   data[3],
					ViewCounter: data[6],
					UpVote:      data[9],
					WriterID:    data[10]})
		case 12:
			if strings.Split(data[11], "|")[0] == "ip" {
				processedResult = append(processedResult,
					ArticleData{URL: data[0],
						Title:       data[1],
						ReplyCount:  "[0]",
						Name:        data[2],
						Timestamp:   data[4],
						ViewCounter: data[7],
						UpVote:      data[10],
						WriterID:    data[11]})
			} else {
				processedResult = append(processedResult,
					ArticleData{URL: data[0],
						Title:       data[1],
						ReplyCount:  data[2],
						Name:        data[3],
						Timestamp:   data[4],
						ViewCounter: data[7],
						UpVote:      data[10],
						WriterID:    data[11]})
			}
		case 13:
			processedResult = append(processedResult,
				ArticleData{URL: data[0],
					Title:       data[1],
					ReplyCount:  data[2],
					Name:        data[3],
					Timestamp:   data[5],
					ViewCounter: data[8],
					UpVote:      data[11],
					WriterID:    data[12]})
		default:
			return nil, errors.New("Data processing error")
		}
	}

	return processedResult, nil
}
