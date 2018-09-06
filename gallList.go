package godc

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

//https://github.com/geeksbaek/goinside/blob/master/request.go
const (
	majorGalleryListAPI = "http://json.dcinside.com/App/gall_name.php"
	minorGalleryListAPI = "http://json.dcinside.com/App/gall_name_sub.php"
)

//FetchMajorGallList download and parse every major gallery list.
//
//모든 메이저 갤러리 목록을 읽어옵니다.
func FetchMajorGallList() ([]GallInfo, error) {
	gallList := fetchURL(majorGalleryListAPI)
	if gallList == nil {
		return nil, errors.New("Page fetch error")
	}
	gallListDoc, err := ioutil.ReadAll(gallList)
	if err != nil {
		return nil, err
	}
	gallListStruct := make([]GallInfo, 0)
	err = json.Unmarshal(gallListDoc, &gallListStruct)
	if err != nil {
		return nil, err
	}
	return gallListStruct, nil
}

//FetchMinorGallList download and parse every major gallery list.
//
//모든 마이너 갤러리 목록을 읽어옵니다.
func FetchMinorGallList() ([]GallInfo, error) {
	gallList := fetchURL(minorGalleryListAPI)
	if gallList == nil {
		return nil, errors.New("Page fetch error")
	}
	gallListDoc, err := ioutil.ReadAll(gallList)
	if err != nil {
		return nil, err
	}
	gallListStruct := make([]GallInfo, 0)
	err = json.Unmarshal(gallListDoc, &gallListStruct)
	if err != nil {
		return nil, err
	}
	return gallListStruct, nil
}
