package mdy

import (
	"fmt"
	"time"
)

const (
	// AppGetURL 获取应用信息 GET
	AppGetURL = "https://api.mingdao.com/v1/open/app/get"

	// AddWorksheetURL 新建工作表 POST
	AddWorksheetURL = "https://api.mingdao.com/v2/open/worksheet/addWorksheet"

	// GetWorksheetInfoURL 获取工作表结构信息 POST
	GetWorksheetInfoURL = "https://api.mingdao.com/v2/open/worksheet/getWorksheetInfo"

	// GetFilterRowsURL 获取列表 POST
	GetFilterRowsURL = "https://api.mingdao.com/v2/open/worksheet/getFilterRows"

	// AddRowURL 新建行记录 POST
	AddRowURL = "https://api.mingdao.com/v2/open/worksheet/addRow"
)

// Response 明道云Response
type Response[T any] struct {
	ErrorMsg  int  `json:"error_msg"`
	ErrorCode int  `json:"error_code"`
	Success   bool `json:"success"`
	Data      T
}

func (r *Response[T]) Ok() bool {
	return r.Success
}

type LocalDateTime time.Time

func (l LocalDateTime) MarshalJSON() ([]byte, error) {
	t := time.Time(l)
	return []byte(fmt.Sprintf("\"%v\"", t.Format("2006-01-02 15:04:05"))), nil
}

func (l LocalDateTime) String() string {
	t := time.Time(l)
	return t.Format("2006-01-02 15:04:05")
}
