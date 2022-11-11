package mdy

import (
	"fmt"
)

type (
	Filter struct {
		ControlId     string        `json:"controlId"` // 字段ID 必填
		DataType      DataType      `json:"dataType"`  // 控件类型编号
		SpliceType    SpliceType    `json:"spliceType,omitempty"`
		FilterType    FilterType    `json:"filterType,omitempty"`
		Value         string        `json:"value,omitempty"`
		Values        []string      `json:"values,omitempty"`
		DateRange     DateRange     `json:"dateRange,omitempty"`
		DateRangeType DateRangeType `json:"dateRangeType,omitempty"`
		MinValue      string        `json:"minValue,omitempty"`
		MaxValue      string        `json:"maxValue,omitempty"`
		IsAsc         bool          `json:"isAsc,omitempty"`        // "是否升序"
		IsGroup       bool          `json:"isGroup,omitempty"`      // "是否升序"
		GroupFilters  []Filter      `json:"groupFilters,omitempty"` // "筛选组列表"
	}
)

// DataType 明道云中Filter筛选器中的数据类型枚举
type DataType int

func (d DataType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", d)), nil
}

const (
	Text        DataType = iota + 2 // 文本 单行、多行
	Phone                           // 电话 手机
	Telephone                       // 电话 座机
	Email                           // 邮箱
	Number                          // 数字
	Certificate                     // 证书
	Amount                          // 金额
	TileSelect                      // 单选 平铺
	Checkboxes                      // 多选
	Select                          // 单选 下拉
	_
	_
	Attachment // 附件
	Date       // 日期: 年-月-日
	DateTime   // 日期: 年-月-日 时:分  16
	_
	_
	Province // 地区: 省
	_
	FreeConnection     // 自由连接
	Subsection         // 分段
	City               // 地区: 省/市
	County             // 地区: 省/市/县
	AmountInWords      // 大写金额
	Member             // 成员
	Department         // 部门
	Level              // 等级
	Associated         // 关联记录
	OtherFields        // 他表字段
	CalculationResults // 公式 计算结果为数字
	TextCombination    // 文本组合
	AutoNumber         // 自动编号
	SubTable           // 子表
	CascadeSelection   // 级联选择
	Checkbox           // 检查框
	Gather             // 汇总
	Formula            // 公式 计算结果为日期
	_
	Location // 定位
	RichText // 富文本
	Sign     // 签名
	_
	_
	Embed                   // 嵌入 45
	Remark DataType = 10010 // 备注
)

// SpliceType 明道云中Filter筛选器中的拼接方式，1:And 2:Or
type SpliceType int

const (
	And SpliceType = iota + 1
	Or
)

func (d SpliceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", d)), nil
}

// FilterType 筛选类型
type FilterType int

func (d FilterType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", d)), nil
}

const (
	Default  FilterType = iota
	Like                // 包含
	Eq                  // 是（等于）
	Start               // 开头为
	End                 // 结尾为
	NContain            // 不包含
	Ne                  // 不是（不等于）
	IsNull              // 为空
	HasValue            // 不为空 8
	_
	_
	Between   // 在范围内 11
	NBetween  // 不在范围内
	Gt        // >
	Gte       // >=
	Lt        // <
	Lte       // <=
	DateEnum  // 日期是
	NDateEnum // 我拥有的 18
	_
	_
	MySelf // 我拥有的 21
	UnRead // 未读
	Sub    // 下属
	RCEq   // 关联控件是
	RCNe   // 关联控件不是
	ArrEq  // 数组等于
	ArrNe  // 数组不等于 27
	_
	_
	_
	DateBetween  // 在范围内 31
	DateNBetween // 不在范围内
	DateGt       // >
	DateGte      // >=
	DateLt       // <
	DateLte      // <= 36
	_
	_
	_
	_
	NormalUser // 常规用户 41
	PortalUser // 外部门户用户
)

const (
	// UserSelf AccountID
	UserSelf     = "user-self"
	UserSub      = "user-sub"
	UserWorkflow = "user-workflow"
	UserApi      = "user-api"
)

// DateRange 日期范围
type DateRange int

func (d DateRange) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", d)), nil
}

const (
	DefaultDateRange DateRange = iota
	Today
	Yesterday
	Tomorrow
	ThisWeek
	LastWeek
	NextWeek
	ThisMonth
	LastMonth
	NextMonth
	LastEnum
	NextEnum
	ThisQuarter
	LastQuarter
	NextQuarter
	ThisYear
	LastYear
	NextYear
	Customize
	_
	_
	Last7Day
	Last14Day
	Last30Day
	Next7Day  DateRange = 31
	Next14Day DateRange = 32
	Next33Day DateRange = 33
)

// DateRangeType 日期范围类型 1：天 2：周 3：月 4：季 5：年
type DateRangeType int

func (d DateRangeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%v", d)), nil
}

const (
	Day DateRangeType = iota + 1
	Week
	Month
	Season
	Year
)
