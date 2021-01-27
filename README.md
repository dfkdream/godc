# godc 
[![Go Report Card](https://goreportcard.com/badge/github.com/dfkdream/godc)](https://goreportcard.com/report/github.com/dfkdream/godc)
[![GoDoc](https://godoc.org/github.com/dfkdream/godc?status.svg)](https://pkg.go.dev/github.com/dfkdream/godc)

디시인사이드 비공식 API

게시글 목록,게시글 내용, 갤러리 목록 읽기 기능만 지원하고 있습니다. 나머지 기능은 추가 예정

## Install

`go get github.com/dfkdream/godc`

## Document

* [godoc](https://pkg.go.dev/github.com/dfkdream/godc)

### Writer
필드 이름 | 자료형 | 설명
---------|--------|------
Name | string | 작성자 이름
Identity | string | 작성자 ID/IP
IsSignedIn | bool | 로그인 여부

### ArticleData

필드 이름 | 자료형 | 설명
---------|--------|------
URL | string | 게시글 URL
Title | string | 게시글 제목
Type | string | 게시글 타입
Tag | string | 게시글 말머리
ReplyCount | string | 댓글 수
Timestamp | string | 작성 시간
ViewCounter | string | 조회수
UpVote | string | 추천 수
Writer | Writer | 작성자

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
Timestamp | string | 작성 시간
ViewCounter | string | 조회수
ReplyCount | string | 댓글 수
Body | string(HTML) | 게시글 내용
UpVote | string | 추천 수
DownVote | string | 비추천 수
Replies | []Reply | 댓글
Writer | Writer | 작성자

#### Reply
필드 이름 | 자료형 | 설명
---------|--------|--------
Type | string | 댓글(reply)/대댓글(re-reply) 구분
Body | string(HTML) | 댓글 내용
Timestamp | string | 작성 시간
Writer | Writer | 작성자

### GallInfo
필드 이름 | 자료형 | 설명
---------|--------|--------
Category | string | 갤러리 카테고리
Name | string | 갤러리 코드
KoName | string | 갤러리 이름
Manager | string | 매니저 ID
SubManager | string( `,` 로 구분) | 부매니저 ID
No | string | 겔러리 번호

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
	dat, err := godc.FetchArticleList("<GalleryID>", 1, false)

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
			data.Timestamp,
			data.ViewCount,
			data.UpVote,
			data.Writer)
	}
}
```
