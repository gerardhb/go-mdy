package mdy

import "github.com/go-resty/resty/v2"

// WorkSheetRequest 代表明道云工作表request
type WorkSheetRequest struct {
	mdy *mdy
	Req *resty.Request
}

// WorkSheetDTO 工作表信息请求DTO
type WorkSheetDTO struct {
	AppKey      string `json:"appKey"`
	Sign        string `json:"sign"`
	WorksheetId string `json:"worksheetId"`
}

// ListDTO 筛选器
type ListDTO struct {
	AppKey       string   `json:"appKey"`      // 必填
	Sign         string   `json:"sign"`        // 签名 必填
	WorksheetId  string   `json:"worksheetId"` // 工作表ID 必填
	ViewId       string   `json:"viewId"`      // "视图ID,必填"
	PageSize     int      `json:"pageSize"`    // 必填
	PageIndex    int      `json:"pageIndex"`   // 必填
	SortId       string   `json:"sortId"`      // "排序字段ID"
	IsAsc        bool     `json:"isAsc"`       // "是否升序"
	Filters      []Filter `json:"filters"`
	NotGetTotal  string   `json:"notGetTotal"`  //"是否不统计总行数以提高性能"
	UseControlId string   `json:"useControlId"` //"是否只返回controlId，默认false"
}

// WorkSheetInfo 代表明道云的每个工作表信息
type WorkSheetInfo struct {
	WorksheetId string `json:"worksheetId"`
	Name        string `json:"name"`
	Views       []struct {
		ViewId string `json:"viewId"`
		Name   string `json:"name"`
	} `json:"views"`
	Controls []struct {
		UserPermission     int           `json:"userPermission"`
		ControlPermissions string        `json:"controlPermissions"`
		CoverCid           string        `json:"coverCid,omitempty"`
		StrDefault         string        `json:"strDefault,omitempty"`
		Size               int           `json:"size"`
		EditAttrs          []interface{} `json:"editAttrs,omitempty"`
		Checked            bool          `json:"checked"`
		ControlId          string        `json:"controlId"`
		Type               int           `json:"type"`
		Attribute          int           `json:"attribute"`
		Hint               string        `json:"hint,omitempty"`
		SourceControlType  int           `json:"sourceControlType"`
		NoticeItem         int           `json:"noticeItem"`
		Required           bool          `json:"required,omitempty"`
		RelationControls   []interface{} `json:"relationControls,omitempty"`
		Unique             bool          `json:"unique"`
		Row                int           `json:"row"`
		Unit               string        `json:"unit,omitempty"`
		EnumDefault2       int           `json:"enumDefault2"`
		Options            []struct {
			Key       string `json:"key"`
			Value     string `json:"value"`
			Index     int    `json:"index"`
			IsDeleted bool   `json:"isDeleted"`
			Color     string `json:"color"`
			Score     int    `json:"score"`
		} `json:"options"`
		AdvancedSetting struct {
			Showtype     string `json:"showtype,omitempty"`
			Max          string `json:"max,omitempty"`
			Itemicon     string `json:"itemicon,omitempty"`
			Analysislink string `json:"analysislink,omitempty"`
			Dismanual    string `json:"dismanual,omitempty"`
			Getinput     string `json:"getinput,omitempty"`
			Getsave      string `json:"getsave,omitempty"`
			Min          string `json:"min,omitempty"`
			Usertype     string `json:"usertype,omitempty"`
			Allowtime    string `json:"allowtime,omitempty"`
			Showformat   string `json:"showformat,omitempty"`
		} `json:"advancedSetting"`
		DeleteAccountId string        `json:"deleteAccountId,omitempty"`
		LastEditTime    string        `json:"lastEditTime"`
		ControlName     string        `json:"controlName"`
		SourceControlId string        `json:"sourceControlId,omitempty"`
		ViewId          string        `json:"viewId,omitempty"`
		Desc            string        `json:"desc,omitempty"`
		FieldPermission string        `json:"fieldPermission,omitempty"`
		Default         string        `json:"default,omitempty"`
		ShowControls    []interface{} `json:"showControls,omitempty"`
		Col             int           `json:"col"`
		EnumDefault     int           `json:"enumDefault"`
		Value           string        `json:"value,omitempty"`
		Disabled        bool          `json:"disabled"`
		Dot             int           `json:"dot"`
		DefaultMen      []interface{} `json:"defaultMen,omitempty"`
		DataSource      string        `json:"dataSource"`
		Half            bool          `json:"half,omitempty"`
		Alias           string        `json:"alias,omitempty"`
		DeleteTime      string        `json:"deleteTime"`
	} `json:"controls"`
}

// GetWorksheetInfo 获取工作表结构信息 POST
func (r *WorkSheetRequest) GetWorksheetInfo(worksheetId string) (*WorkSheetInfo, error) {
	response, err := r.Req.SetResult(&Response[WorkSheetInfo]{}).SetBody(&WorkSheetDTO{
		AppKey:      r.mdy.appKey,
		Sign:        r.mdy.sign,
		WorksheetId: worksheetId,
	}).Post(GetWorksheetInfoURL)
	if err != nil {
		return nil, err
	}
	result := response.Result().(*Response[WorkSheetInfo])
	return &result.Data, nil
}

// GetFilterRows 获取列表 POST
func (r *WorkSheetRequest) GetFilterRows(worksheetId string) (*WorkSheetInfo, error) {
	response, err := r.Req.SetResult(&Response[WorkSheetInfo]{}).SetBody(&WorkSheetDTO{
		AppKey:      r.mdy.appKey,
		Sign:        r.mdy.sign,
		WorksheetId: worksheetId,
	}).Post(GetWorksheetInfoURL)
	if err != nil {
		return nil, err
	}
	result := response.Result().(*Response[WorkSheetInfo])
	return &result.Data, nil
}
