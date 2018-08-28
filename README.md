# godc 
[![Go Report Card](https://goreportcard.com/badge/github.com/dfkdream/godc)](https://goreportcard.com/report/github.com/dfkdream/godc)
[![GoDoc](https://godoc.org/github.com/dfkdream/godc?status.svg)](https://godoc.org/github.com/dfkdream/godc)

디시인사이드 비공식 API

게시글 목록,게시글 내용 읽기 기능만 지원하고 있습니다. 나머지 기능은 추가 예정

## Install

`go get github.com/dfkdream/godc`

## Document

* [godoc](https://godoc.org/github.com/dfkdream/godc)

### ArticleData

필드 이름 | 자료형 | 설명
---------|--------|------
URL | string | 게시글 URL
Title | string | 게시글 제목
Type | string | 게시글 타입
ReplyCount | string | 댓글 수
Name | string | 작성자 이름
Timestamp | string | 작성 시간
ViewCounter | string | 조회수
UpVote | string | 추천 수
WriterID | string( `\|` 로 구분) | 작성자 ID/IP

#### ArticleData.Type

string | 설명
-----|-----
`ico_pic ico_t` | 텍스트만
`ico_pic ico_p_y` | 이미지 포함
`ico_pic ico_mv` | 동영상 포함
`ico_pic ico_t_c` | 텍스트만, 개념글
`ico_pic ico_p_c` | 이미지 포함, 개념글

### ArticleBody
필드 이름 | 자료형 | 설명
---------|--------|-------
Title | string | 게시글 제목
Name | string | 작성자 이름
IP | string | 작성자 IP(유동일 경우만)
Timestamp | string | 작성 시간
ViewCounter | string | 조회수
Body | string(HTML) | 게시글 내용
UpVote | string | 추천 수
DownVote | string | 비추천 수
Replies | []Reply | 댓글

#### Reply
필드 이름 | 자료형 | 설명
---------|--------|--------
URL | string | 갤로그 주소
Name | string | 작성자 이름
IP | string | 작성자 IP(유동일 경우만)
Body | string(HTML) | 댓글 내용
Timestamp | string | 작성 시간

## Example Code

GalleryID 갤러리의 게시글 목록 1페이지를 읽어옵니다.
```Go
package main

import (
	"fmt"
	"log"

	"github.com/dfkdream/godc"
)

func main() {
	dat, err := godc.FetchArticleList("yurucam", 1)

	if err != nil {
		log.Fatal(err)
	}

	for index, data := range dat {
		fmt.Printf("=============article%d==============\n", index)
		fmt.Printf("URL: %s\nTitle: %s\nType: %s\nReplyCount: %s\nName : %s\nTimestamp : %s\nViewCounter : %s\nUpVote : %s\nWriterID : %s\n",
			data.URL,
			data.Title,
			data.Type,
			data.ReplyCount,
			data.Name,
			data.Timestamp,
			data.ViewCounter,
			data.UpVote,
			data.WriterID)
	}
}
```
