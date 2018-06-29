package godc

import (
	"io"
	"net/http"
	"strconv"
	"time"
)

func fetchPage(gallCode string, page int) io.ReadCloser {
	client := &http.Client{
		Timeout: time.Second * 30,
	}
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
