package mdy

import (
	jsoniter "github.com/json-iterator/go"
	"testing"
	"time"
)

// 对象转 controls
type ObjectDTO struct {
	//Lianxiren1   string        `json:"lianxiren1" mdy:"ctlId:636479496f7adb33b5b2a1f9"` // 联系人
	Qtzdynr      string        `mdy:"ctlId:636479496f7adb33b5b2a1fa"` // 其他字段与内容
	Fujian       string        `mdy:"ctlId:636479496f7adb33b5b2a1fb;valueType:1;editType:1"`
	Xslysj       LocalDateTime `mdy:"ctlId:636479496f7adb33b5b2a1fd"`            // 线索来源时间 日期和时间
	Xslylb       []string      `mdy:"ctlId:636479496f7adb33b5b2a1fe"`            // 线索来源类别
	Xslyr        string        `mdy:"ctlId:636479496f7adb33b5b2a200"`            // 线索来源人
	Xslyqd       string        `mdy:"ctlId:636479496f7adb33b5b2a1ff;default:展会"` // 线索来源渠道
	Xszqdgtqk    string        `mdy:"ctlId:636479496f7adb33b5b2a20a"`            // 线索之前的沟通情况
	Syhzt_x_     string        `mdy:"ctlId:636479496f7adb33b5b2a213"`            // 森友会状态(新）
	Jqqk         string        `mdy:"ctlId:636479496f7adb33b5b2a214"`            // 进群情况
	Qunmingcheng string        `mdy:"ctlId:636479496f7adb33b5b2a210"`            // 群名称 表关联
	Sfcjxczb     string        `mdy:"ctlId:636479496f7adb33b5b2a215"`            // 是否参加现场直播
	Sffssyhzl    string        `mdy:"ctlId:636479496f7adb33b5b2a216"`            // 是否发送森友会资料
	Khyx         string        `mdy:"ctlId:636479496f7adb33b5b2a217"`            // 客户意向
	Khtd         string        `mdy:"ctlId:636479496f7adb33b5b2a218"`            // 客户态度
	Lianxiren3   string        `mdy:"ctlId:636479496f7adb33b5b2a208"`            // 联系人 表关联
	Qiye2        string        `mdy:"ctlId:636479496f7adb33b5b2a20c"`            // 企业 表关联
}

type Enterprise struct {
}

type Contact struct {
}

func TestAddRow(t *testing.T) {
	obj := ObjectDTO{
		Qtzdynr:      "备注",
		Fujian:       "https://static.thingclub.com/thingyouwe-oss/20221103/09/1394ffb0-2ba8-46da-8a32-a610297e5816.png",
		Xslysj:       LocalDateTime(time.Now()),
		Xslylb:       []string{"与优锘/森友会合作（会员线索）"},
		Xslyqd:       "",
		Xslyr:        "长沙",
		Xszqdgtqk:    "线索之前的沟通情况叭叭叭八",
		Syhzt_x_:     "未注册",
		Jqqk:         "未邀请",
		Qunmingcheng: "",
		Sfcjxczb:     "是",
		Sffssyhzl:    "否",
		Khyx:         "了解一下",
		Khtd:         "不耐烦",
		Lianxiren3:   "",
		Qiye2:        "",
	}
	controls, err := ToControls(&obj)
	if err != nil {
		t.Error(err)
	}

	newMdy := NewMdy()
	dto := &RowDTO{
		WorksheetId: "636479496f7adb33b5b2a0fd",
		Controls:    controls,
	}
	row, err := newMdy.WorkSheetReq().AddRow(dto)
	if err != nil {
		t.Error(err)
	}

	t.Log(row)

	//t.Logf("type :%T", controls)
	//t.Log(controls)
	//
	//indent, err := json.MarshalIndent(controls, "", "    ")
	//if err != nil {
	//	t.Error(err)
	//}
	//t.Log(string(indent))
}

func TestWorkSheetRequest_GetFilterRows(t *testing.T) {
	newMdy := NewMdy()
	dto := &FilterRowDTO{
		BaseDTO:     BaseDTO{},
		WorksheetId: "622847896b1004fcfbec4d4d",
		ViewId:      "622847896b1004fcfbec4d51",
		PageSize:    10,
		PageIndex:   1,
		SortId:      "",
		IsAsc:       false,
		Filters: []Filter{{
			ControlId:  "62284bb8dc31f4a577861f59",
			DataType:   Text,
			SpliceType: And,
			FilterType: Eq,
			Value:      "圆融云海（山东）软件技术有限公司",
			//Values:        nil,
			//DateRange:     0,
			//DateRangeType: 0,
			//MinValue:      "",
			//MaxValue:      "",
			//IsAsc:         false,
			//IsGroup:       false,
			//GroupFilters:  nil,
		}},
		NotGetTotal:  true, //
		UseControlId: true,
	}
	rows, err := newMdy.WorkSheetReq().GetFilterRows(dto)
	if err != nil {
		t.Error(err)
	}
	indent, err := jsoniter.MarshalIndent(rows, "", "    ")
	if err != nil {
		t.Error(err)
	}
	t.Log(string(indent))
}
