//Package godc reads list of posts from the DCInside Gallery.
//
//디시인사이드 갤러리 게시글 목록을 읽어오는 패키지
package godc

import (
	"io"
	"net/http"
	"strconv"

	"golang.org/x/net/html"
)

//ArticleData contains post URL, post title, username.
//
//게시글 주소(URL), 게시글 제목(Title), 닉네임(고닉/유동 구분 안됨)(Name)
type ArticleData struct {
	URL   string
	Title string
	Name  string
}

func fetchPage(gallCode string, page int) io.ReadCloser {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://m.dcinside.com/list.php?id="+gallCode+"&page="+strconv.Itoa(page), nil)
	if err != nil {
		return nil
	}
	req.Header.Set("Host", "m.dcinside.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.98 Mobile Safari/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	req.Header.Set("Accept-Language", "ko-KR,ko;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "utf-8")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	return resp.Body
}

//FetchAndParsePage reads specific page of post list.
//
//게시글 목록의 지정된 페이지를 읽어옵니다
func FetchAndParsePage(gallID string, page int) ([]ArticleData, error) {
	dcpg := fetchPage(gallID, page)
	doc, err := html.Parse(dcpg)
	if err != nil {
		return nil, err
	}
	result := make([]ArticleData, 0)
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "ul" && len(n.Attr) > 0 && n.Attr[0].Val == "list_best" {
			for li := n.FirstChild; li != nil; li = li.NextSibling {
				if li.Type == html.ElementNode && li.Data == "li" && len(li.Attr) == 0 {
					for span := li.FirstChild; span != nil; span = span.NextSibling {
						if span.Type == html.ElementNode && span.Data == "span" {
							for a := span.FirstChild; a != nil; a = a.NextSibling {
								if a.Type == html.ElementNode && a.Data == "a" {
									articleURL := a.Attr[0].Val
									articleTitle := ""
									articleName := ""
									for spanData := a.FirstChild; spanData != nil; spanData = spanData.NextSibling {
										if spanData.Type == html.ElementNode && spanData.Data == "span" {
											switch spanData.Attr[0].Val {
											case "title":
												for titleData := spanData.FirstChild; titleData != nil; titleData = titleData.NextSibling {
													if titleData.Type == html.ElementNode && titleData.Data == "span" && len(titleData.Attr) > 0 && titleData.Attr[0].Val == "txt" {
														articleTitle = titleData.FirstChild.Data
													}
												}
											case "info":
												for infoData := spanData.FirstChild; infoData != nil; infoData = infoData.NextSibling {
													if infoData.Type == html.ElementNode && infoData.Data == "span" && len(infoData.Attr) > 0 && infoData.Attr[0].Val == "name" {
														articleName = infoData.FirstChild.Data
													}
												}
											default:
												continue
											}
										}
									}
									result = append(result, ArticleData{articleURL, articleTitle, articleName})
								}
							}
						}
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return result, nil
}
