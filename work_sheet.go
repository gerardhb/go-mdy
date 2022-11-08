package mdy

import "github.com/go-resty/resty/v2"

// WorkSheetRequest 代表明道云工作表request
type WorkSheetRequest struct {
	mdy *mdy
	Req *resty.Request
}

type BaseDTO struct {
	AppKey string `json:"appKey"`
	Sign   string `json:"sign"`
}

func (b *BaseDTO) SetSign(sign string) {
	b.Sign = sign
}

func (b *BaseDTO) SetAppKey(appKey string) {
	b.AppKey = appKey
}

// WorkSheetDTO 工作表信息请求DTO
type WorkSheetDTO struct {
	BaseDTO
	WorksheetId string `json:"worksheetId"`
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

func newWorkSheetDTO(appKey, sign, worksheetId string) *WorkSheetDTO {
	return &WorkSheetDTO{
		BaseDTO: BaseDTO{
			AppKey: appKey,
			Sign:   sign,
		},
		WorksheetId: worksheetId,
	}
}

// GetWorksheetInfo 获取工作表结构信息 POST
func (r *WorkSheetRequest) GetWorksheetInfo(worksheetId string) (*WorkSheetInfo, error) {
	dto := &WorkSheetDTO{
		WorksheetId: worksheetId,
	}
	r.Set(dto)
	response, err := r.Req.SetResult(&Response[WorkSheetInfo]{}).
		SetBody(dto).Post(GetWorksheetInfoURL)
	if err != nil {
		return nil, err
	}
	result := response.Result().(*Response[WorkSheetInfo])
	return &result.Data, nil
}

// GetFilterRows 获取列表 POST
func (r *WorkSheetRequest) GetFilterRows(filter *FilterRowDTO) (*PageRow, error) {
	r.Set(filter)
	response, err := r.Req.SetResult(&Response[PageRow]{}).SetBody(filter).Post(GetFilterRowsURL)
	if err != nil {
		return nil, err
	}
	result := response.Result().(*Response[PageRow])
	return &result.Data, nil
}

// AddRow 新建行记录 POST
// return 记录的id
func (r *WorkSheetRequest) AddRow(row *RowDTO) (string, error) {
	r.Set(row)
	response, err := r.Req.SetResult(&Response[string]{}).SetBody(row).Post(AddRowURL)
	if err != nil {
		return "", err
	}
	result := response.Result().(*Response[string])
	return result.Data, nil
}

// EditRow 新建行记录 POST
// return 记录的id
func (r *WorkSheetRequest) EditRow(row *RowDTO) (string, error) {
	r.Set(row)
	response, err := r.Req.SetResult(&Response[string]{}).SetBody(row).Post(EditRowURL)
	if err != nil {
		return "", err
	}
	result := response.Result().(*Response[string])
	return result.Data, nil
}

func (r *WorkSheetRequest) Set(info any) {
	if sign, ok := info.(Signer); ok {
		sign.SetSign(r.mdy.GetSign())
		sign.SetAppKey(r.mdy.GetAppKey())
	}
}
