package godc

import "testing"

func isGallInfoOK(v GallInfo) bool {
	if v.Name == "" ||
		v.Category == "" ||
		v.KoName == "" ||
		v.No == "" {
		return false
	}
	return true
}

func TestFetchMajorGallList(t *testing.T) {
	l, err := FetchMajorGallList()
	if err != nil {
		t.Error(err)
	}

	if len(l) < 1 {
		t.Error("expected len(l)>0 but got 0")
	}

	for i, v := range l {
		if !isGallInfoOK(v) {
			t.Errorf("missing required field: %d %+v", i, v)
		}
	}
}

func TestFetchMinorGallList(t *testing.T) {
	l, err := FetchMajorGallList()
	if err != nil {
		t.Error(err)
	}

	if len(l) < 1 {
		t.Error("expected len(l)>0 but got 0")
	}

	for i, v := range l {
		if !isGallInfoOK(v) {
			t.Errorf("missing required field: %d %+v", i, v)
		}
	}
}
