package godc

import (
	"bytes"
	"errors"
	"io"
	"regexp"
	"strconv"
	"strings"

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
	doc, err := html.Parse(dcArticle)
	if err != nil {
		return nil, errors.New("html parse error")
	}

	result := ArticleBody{}

	//Body 처리
	var gallContent *html.Node
	var searchGallContent func(*html.Node)
	searchGallContent = func(n *html.Node) {
		if n.Type == html.ElementNode && len(n.Attr) > 0 && n.Attr[0].Val == "gall_content" {
			gallContent = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchGallContent(c)
		}
	}
	searchGallContent(doc)

	var searchTitle func(*html.Node)
	searchTitle = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "tit_view" {
			tmpResult := ""
			for d := n.FirstChild; d != nil; d = d.NextSibling {
				if d != nil && d.Data != "img" {
					tmpResult += renderNode(d)
				}
			}
			result.Title = strings.TrimSpace(tmpResult)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchTitle(c)
		}
	}
	searchTitle(gallContent)

	var infoHeader *html.Node
	var searchInfoHeader func(*html.Node)
	searchInfoHeader = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "info_edit" {
			infoHeader = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchInfoHeader(c)
		}
	}
	searchInfoHeader(gallContent)

	var searchName func(*html.Node)
	vaildTimestamp := regexp.MustCompile(`\d+\.\d+\.\d+ \d+\:\d+$`)
	searchName = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) == 0 && n.FirstChild != nil {
			if vaildTimestamp.MatchString(n.FirstChild.Data) {
				result.Timestamp = n.FirstChild.Data
			} else {
				tmpResult := ""
				for d := n.FirstChild; d != nil; d = d.NextSibling {
					if d != nil && d.Data != "img" {
						tmpResult += renderNode(d)
					}
				}
				result.Name = strings.TrimSpace(tmpResult)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchName(c)
		}
	}
	searchName(infoHeader)

	var searchViewCounter func(*html.Node)
	searchViewCounter = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) == 1 && n.Attr[0].Val == "num" && n.FirstChild != nil {
			result.ViewCounter = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchViewCounter(c)
		}
	}
	searchViewCounter(infoHeader)

	var searchReplyCount func(*html.Node)
	searchReplyCount = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 1 && n.Attr[1].Val == "comment_dirc" && n.FirstChild != nil {
			result.ReplyCount = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchReplyCount(c)
		}
	}
	searchReplyCount(infoHeader)

	var searchBody func(*html.Node)
	searchBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && len(n.Attr) > 0 && n.Attr[0].Val == "view_main" {
			result.Body = removeScript(strings.TrimSpace(renderNode(n)))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchBody(c)
		}
	}
	searchBody(gallContent)

	var searchUpVote func(*html.Node)
	searchUpVote = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "recomm_btn" && n.FirstChild != nil {
			result.UpVote = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchUpVote(c)
		}
	}
	searchUpVote(gallContent)

	var searchDownVote func(*html.Node)
	searchDownVote = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "nonrecomm_btn" && n.FirstChild != nil {
			result.DownVote = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchDownVote(c)
		}
	}
	searchDownVote(gallContent)

	var boxShare *html.Node
	var searchBoxShare func(*html.Node)
	searchBoxShare = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "div" && len(n.Attr) > 0 && n.Attr[0].Val == "box_share" {
			boxShare = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchBoxShare(c)
		}
	}
	searchBoxShare(gallContent)

	var searchIP func(*html.Node)
	searchIP = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "ip" && n.FirstChild != nil {
			result.IP = n.FirstChild.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchIP(c)
		}
	}
	searchIP(boxShare)

	var searchGallogURL func(*html.Node)
	searchGallogURL = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" && len(n.Attr) > 1 && n.Attr[1].Val == "btn btn_gall" {
			result.GallogURL = n.Attr[0].Val
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchGallogURL(c)
		}
	}
	searchGallogURL(boxShare)

	if commentNo, err := strconv.Atoi(result.ReplyCount); err == nil && commentNo == 0 {
		return &result, nil
	}
	//Comment 처리
	gallComment, err := html.Parse(fetchAllReply(URL))
	if err != nil {
		return nil, err
	}
	/*	var searchGallComment func(*html.Node)
		searchGallComment = func(n *html.Node) {
			if n.Type == html.ElementNode && len(n.Attr) > 0 && n.Attr[0].Val == "wrap_list" {
				gallComment = n
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				searchGallComment(c)
			}
		}
		searchGallComment(doc)*/

	comments := make([]*html.Node, 0)
	var sepComments func(*html.Node)
	sepComments = func(n *html.Node) {
		var vaildComment = regexp.MustCompile(`comment_cnt_[0-9]*$`)
		if n != nil && n.Type == html.ElementNode && n.Data == "li" && len(n.Attr) > 0 && vaildComment.MatchString(n.Attr[0].Val) {
			comments = append(comments, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			sepComments(c)
		}
	}
	sepComments(gallComment)

	for _, commentNode := range comments {
		parsedComment := Reply{}
		var parseName func(*html.Node)
		parseName = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) == 0 && n.FirstChild != nil {
				parsedComment.Name = n.FirstChild.Data[1:]
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseName(c)
			}
		}
		parseName(commentNode)

		var parseIPName func(*html.Node)
		parseIPName = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "id" && n.FirstChild != nil {
				rCombined := ""
				for d := n.FirstChild; d != nil; d = d.NextSibling {
					rCombined += renderNode(d)
				}
				parsedComment.Name = rCombined[1:][:len(rCombined)-2]
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseIPName(c)
			}
		}
		parseIPName(commentNode)

		var parseURL func(*html.Node)
		parseURL = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" && len(n.Attr) > 0 && n.Attr[1].Val == "id" && n.FirstChild != nil {
				parsedComment.URL = n.Attr[0].Val
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseURL(c)
			}
		}
		parseURL(commentNode)

		var parseBody func(*html.Node)
		parseBody = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "txt" && n.FirstChild != nil {
				parsedComment.Body = removeScript(strings.TrimSpace(renderNode(n.FirstChild)))
			}
			if parsedComment.Body == "" {
				if n.Type == html.ElementNode && n.Data == "img" {
					parsedComment.Body = removeScript(strings.TrimSpace(renderNode(n)))
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseBody(c)
			}
		}
		parseBody(commentNode)

		var parseTimestamp func(*html.Node)
		parseTimestamp = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "date" && n.FirstChild != nil {
				parsedComment.Timestamp = n.FirstChild.Data
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseTimestamp(c)
			}
		}
		parseTimestamp(commentNode)

		var parseIP func(*html.Node)
		parseIP = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "span" && len(n.Attr) > 0 && n.Attr[0].Val == "ip" && n.FirstChild != nil {
				parsedComment.IP = n.FirstChild.Data
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parseIP(c)
			}
		}
		parseIP(commentNode)

		result.Replies = append(result.Replies, parsedComment)
	}

	return &result, nil

}
