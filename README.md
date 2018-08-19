# Go-DCApi
[![Go Report Card](https://goreportcard.com/badge/github.com/dfkdream/Go-DCApi)](https://goreportcard.com/report/github.com/dfkdream/Go-DCApi)
[![GoDoc](https://godoc.org/github.com/dfkdream/Go-DCApi?status.svg)](https://godoc.org/github.com/dfkdream/Go-DCApi)

디시인사이드 비공식 API

게시글 목록 읽기 기능만 지원하고 있습니다. 나머지 기능은 추가 예정

## Install

`go get github.com/dfkdream/Go-DCApi`

## Document

* [godoc](https://godoc.org/github.com/dfkdream/Go-DCApi)

### ArticleData

필드 이름 | 자료형 | 설명
---------|--------|------
URL | string | 게시글 URL
Title | string | 게시글 제목
ReplyCount | string | 댓글 수
Name | string | 작성자 이름
Timestamp | string | 작성 시간
ViewCounter | string | 조회수
UpVote | string | 추천 수
WriterID | string( `\|` 로 구분) | 작성자 ID/IP

## Example Code

GalleryID 갤러리의 게시글 목록 1페이지를 읽어옵니다.
```Go
package main

import (
    "fmt"
    "log"

    "github.com/dfkdream/Go-DCApi"
)

func main() {
    dat, err := godc.FetchAndParsePage("galleryID", 1)

    if err != nil {
        log.Fatal(err)
    }

    for index, data := range dat {
        fmt.Printf("=============article%d==============\n", index)
        fmt.Printf("URL: %s\nTitle: %s\nReplyCount: %s\nName : %s\nTimestamp : %s\nViewCounter : %s\nUpVote : %s\nWriterID : %s\n",
            data.URL,
            data.Title,
            data.ReplyCount,
            data.Name,
            data.Timestamp,
            data.ViewCounter,
            data.UpVote,
            data.WriterID)
    }
}
```
