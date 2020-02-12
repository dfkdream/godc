package godc

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchReply(URL string, page int) io.ReadCloser {
	var gallCode string
	var articleNo string

	uSpl := strings.Split(URL, "/")
	gallCode = uSpl[len(uSpl)-2]
	articleNo = uSpl[len(uSpl)-1]

	client := &http.Client{
		//Timeout: time.Second * 30,
	}

	form := url.Values{}
	form.Add("id", gallCode)
	form.Add("no", articleNo)
	form.Add("cpage", strconv.Itoa(page))
	form.Add("managerskill", "")
	form.Add("del_scope", "")
	form.Add("csort", "")

	reqURL := "https://m.dcinside.com/ajax/response-comment"
	req, err := http.NewRequest("POST", reqURL, strings.NewReader(form.Encode()))
	if err != nil {
		return nil
	}

	req.Header.Set("Host", "m.dcinside.com")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 7.0; PLUS Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.98 Mobile Safari/537.36")
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ko-KR,ko;q=0.8,en-US;q=0.5,en;q=0.3")
	req.Header.Set("Accept-Encoding", "utf-8")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("X-Requested-With", "XMLHttpRequest")
	req.Header.Set("Content-Length", strconv.Itoa(len(form.Encode())))
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Upgrade-Insecure-Requests", "1")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Referer", URL)

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	return resp.Body
}

func parseReply(URL string, page int) ([]Reply, bool, error) {
	rReply := fetchReply(URL, page)
	if rReply == nil {
		return nil, false, errors.New("Reply fetch error")
	}

	qRep, err := goquery.NewDocumentFromReader(rReply)
	if err != nil {
		return nil, false, err
	}

	comments := make([]Reply, 0)

	qRep.Find("li").Each(func(i int, s *goquery.Selection) {
		if _, exist := s.Attr("id"); !exist {
			return
		}
		var comment Reply
		if v, _ := s.Attr("class"); v == "comment " {
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
		comment.ID = s.Find("span.blockCommentId").AttrOr("data-info", "")
		comment.IP = s.Find("span.ip").Text()
		comment.Timestamp = s.Find("span.date").Text()
		commentbHTML, _ := s.Find("p.txt").Html()
		comment.Body = removeScript(strings.TrimSpace(commentbHTML))
		comments = append(comments, comment)
	})

	if hr, exist := qRep.Find("a.next").Attr("href"); exist && hr != "javascript:;" {
		return comments, true, nil
	}
	return comments, false, nil

}

func fetchAllReply(URL string) ([]Reply, error) {
	result := make([]Reply, 0)
	page := 1
	for {
		replies, isNextAvailable, err := parseReply(URL, page)
		if err != nil {
			return result, err
		}
		result = append(replies, result...)
		if !isNextAvailable {
			break
		}
		page++
	}
	return result, nil
}
