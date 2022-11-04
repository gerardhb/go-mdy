package mdy

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"testing"
)

func TestWorkSheetRequest_GetWorksheetInfo(t *testing.T) {
	newMdy := NewMdy()

	data, err := newMdy.WorkSheetReq().GetWorksheetInfo("622847896b1004fcfbec4d4d")
	if err != nil {
		t.Error(err)
	}
	indent, err := jsoniter.MarshalIndent(data, "", "    ")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(indent))

}
