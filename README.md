# Go-DCApi
[![Go Report Card](https://goreportcard.com/badge/github.com/dfkdream/Go-DCApi)](https://goreportcard.com/report/github.com/dfkdream/Go-DCApi)
[![GoDoc](https://godoc.org/github.com/dfkdream/Go-DCApi?status.svg)](https://godoc.org/github.com/dfkdream/Go-DCApi)

디시인사이드 갤러리 게시글 목록을 읽어오는 패키지

## Install

`go get github.com/dfkdream/Go-DCApi`

## Document

* [godoc](https://godoc.org/github.com/dfkdream/Go-DCApi)

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

	for _, doc := range dat {
		fmt.Println("=============page1================")
		fmt.Println(doc.URL)
		fmt.Println(doc.Title)
		fmt.Println(doc.Name)
	}
}
```
