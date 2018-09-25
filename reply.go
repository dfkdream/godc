package godc

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func fetchReply(URL string, page int) io.ReadCloser {
	reqTmp, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil
	}
	gallCode := reqTmp.URL.Query().Get("id")
	articleNo := reqTmp.URL.Query().Get("no")
	if gallCode == "" || articleNo == "" {
		return nil
	}

	client := &http.Client{
		Timeout: time.Second * 30,
	}
	reqURL := fmt.Sprintf("http://m.dcinside.com/comment_more_new.php?id=%s&no=%s&com_page=%d", gallCode, articleNo, page)
	req, err := http.NewRequest("GET", reqURL, nil)
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
	req.Header.Set("Referer", URL)

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	return resp.Body
}

func fetchAllReply(URL string) *strings.Reader {
	result := ""
	page := 1
	for {
		docReader := fetchReply(URL, page)
		resTmp, err := ioutil.ReadAll(docReader)
		if err != nil {
			return nil
		}
		resPart := string(resTmp)
		result = resPart + result
		if !strings.Contains(resPart, `button type="button" class="btn_page mr" onclick="comment_more(`) {
			break
		}
		page++
	}
	return strings.NewReader(result)
}
