package mdy

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
	"strings"
)

type RowParams map[string]any

type ControlsGeneric interface {
	[]*ControlDTO | [][]*ControlDTO
}

type (

	// FilterRowDTO 列表获取筛选器
	FilterRowDTO struct {
		BaseDTO
		WorksheetId  string   `json:"worksheetId"` // 工作表ID 必填
		ViewId       string   `json:"viewId"`      // "视图ID,必填"
		PageSize     int      `json:"pageSize"`    // 必填
		PageIndex    int      `json:"pageIndex"`   // 必填
		SortId       string   `json:"sortId"`      // "排序字段ID"
		IsAsc        bool     `json:"isAsc"`       // "是否升序"
		Filters      []Filter `json:"filters"`
		NotGetTotal  bool     `json:"notGetTotal"`  //"是否不统计总行数以提高性能"
		UseControlId bool     `json:"useControlId"` //"是否只返回controlId，默认false"
	}

	// RowDTO 新增行请求DTO
	RowDTO struct {
		BaseDTO
		WorksheetId     string          `json:"worksheetId"`
		RowId           string          `json:"rowId,omitempty"`
		TriggerWorkflow bool            `json:"triggerWorkflow,omitempty"`
		Controls        []*ControlDTO   `json:"controls,omitempty"`
		Rows            [][]*ControlDTO `json:"rows,omitempty"`
	}

	ControlDTO struct {
		ControlId    string `json:"controlId"`
		Value        string `json:"value"`
		EditType     int64  `json:"editType"`
		ValueType    int64  `json:"valueType"`
		ControlFiles []struct {
			BaseFile string `json:"baseFile"`
			FileName string `json:"fileName"`
		} `json:"controlFiles,omitempty"`
	}
)

// fields of ControlDTO index
const (
	controlIdIndex = 0
	valueIndex     = 1
	editTypeIndex  = 2
	valueTypeIndex = 3
)

// IdBody 一些附件信息
type IdBody struct {
	AccountId string `json:"accountId"`
	Fullname  string `json:"fullname"`
	Avatar    string `json:"avatar"`
	Status    int    `json:"status"`
}

// MapToIdBody map convert to *IdBody
func MapToIdBody(data map[string]any) *IdBody {
	if data == nil {
		return nil
	}
	return &IdBody{
		AccountId: data["accountId"].(string),
		Fullname:  data["fullname"].(string),
		Avatar:    data["avatar"].(string),
		Status:    int(data["status"].(float64)),
	}
}

type Row struct {
	RowParams
	Ctime              string  `json:"ctime"`
	Utime              string  `json:"utime"`
	Wfftime            string  `json:"wfftime"`
	AutoId             int64   `json:"autoid"`
	AllowDelete        bool    `json:"allowdelete"`
	ControlPermissions string  `json:"controlpermissions"`
	Ownerid            *IdBody `json:"ownerid"`
	Uaid               *IdBody `json:"uaid"`
	Caid               *IdBody `json:"caid"`
	Rowid              string  `json:"rowid"`
}

func (r *Row) UnmarshalJSON(bytes []byte) error {
	if r.RowParams == nil {
		r.RowParams = make(map[string]any, 10)
	}

	err := jsoniter.Unmarshal(bytes, &r.RowParams)
	if err != nil {
		return err
	}

	// 固定参数转换
	if ctime, ok := r.RowParams["ctime"]; ok {
		r.Ctime = ctime.(string)
		delete(r.RowParams, "ctime")
	}

	if utime, ok := r.RowParams["utime"]; ok {
		r.Utime = utime.(string)
		delete(r.RowParams, "utime")
	}

	if wfftime, ok := r.RowParams["wfftime"]; ok {
		r.Wfftime = wfftime.(string)
		delete(r.RowParams, "wfftime")
	}

	if autoid, ok := r.RowParams["autoid"]; ok {
		if id, okFloat := autoid.(float64); okFloat {
			r.AutoId = int64(id)
			delete(r.RowParams, "autoid")
		}
	}

	if allowdelete, ok := r.RowParams["allowdelete"]; ok {
		r.AllowDelete = allowdelete.(bool)
		delete(r.RowParams, "allowdelete")
	}

	if controlpermissions, ok := r.RowParams["controlpermissions"]; ok {
		r.ControlPermissions = controlpermissions.(string)
		delete(r.RowParams, "controlpermissions")
	}

	if rowid, ok := r.RowParams["rowid"]; ok {
		r.Rowid = rowid.(string)
		delete(r.RowParams, "rowid")
	}

	if ownerid, ok := r.RowParams["ownerid"]; ok {
		if owneridObj, ok := ownerid.(map[string]any); ok {
			r.Ownerid = MapToIdBody(owneridObj)
			delete(r.RowParams, "ownerid")
		}
	}

	if uaid, ok := r.RowParams["uaid"]; ok {
		if uaidObj, ok := uaid.(map[string]any); ok {
			r.Uaid = MapToIdBody(uaidObj)
			delete(r.RowParams, "uaid")
		}
	}

	if caid, ok := r.RowParams["caid"]; ok {
		if caidObj, ok := caid.(map[string]any); ok {
			r.Caid = MapToIdBody(caidObj)
			delete(r.RowParams, "caid")
		}
	}

	return nil
}

type PageRow struct {
	Rows  []Row `json:"rows"`
	Total int64 `json:"total"`
}

// Empty length of the Rows is zero
func (p *PageRow) Empty() bool {
	return len(p.Rows) == 0
}

// ToControls 把实际对象数据转成明道云提交的Control
func ToControls(data any) (result []*ControlDTO, err error) {
	defer func() {
		if errRec := recover(); err != nil {
			err = fmt.Errorf("%v", errRec)
		}
	}()

	dataValue := reflect.ValueOf(data)
	if !dataValue.IsValid() {
		return nil, fmt.Errorf("toControls prarm is invalid: %v", dataValue)
	}

	// 获取实际类型
	trueValue := reflect.Indirect(dataValue)
	// 构建返回值
	ctlType := reflect.TypeOf(ControlDTO{})
	slice := reflect.MakeSlice(reflect.TypeOf([]*ControlDTO{}), 0, trueValue.NumField())

	for i := 0; i < trueValue.NumField(); i++ {
		structField := trueValue.Type().Field(i)
		val := trueValue.Field(i)

		mdyTags := structField.Tag.Get("mdy")
		if mdyTags == "" || mdyTags == "-" {
			continue
		}
		// 空值不处理
		if val.IsZero() && !strings.Contains(mdyTags, "default:") {
			continue
		}

		// ControlDTO的指针
		controlPtr := reflect.New(ctlType)
		// ControlDTO的实际Value
		control := controlPtr.Elem()
		if err := tagsHandle(mdyTags, control); err != nil {
			return nil, err
		}

		controlValueField := control.Field(valueIndex)

		err := setControlValue(controlValueField, val, structField)
		if err != nil {
			return nil, err
		}

		slice = reflect.Append(slice, controlPtr)
	}

	return slice.Interface().([]*ControlDTO), nil
}

// tagsHandle 对象标签处理,把tag中的值赋值到control对象对应的字段上
func tagsHandle(mdyTags string, control reflect.Value) error {
	mdyTagStr := strings.Split(mdyTags, ";")
	for i := range mdyTagStr {
		keyValues := strings.Split(mdyTagStr[i], ":")
		if len(keyValues) == 2 {
			key := keyValues[0]
			value := keyValues[1]
			if key == "ctlId" {
				controlIdField := control.Field(controlIdIndex)
				if controlIdField.IsValid() {
					controlIdField.SetString(value)
				}
			} else if key == "valueType" {
				val, err := strconv.Atoi(value)
				if err != nil {
					return fmt.Errorf("valueType error converting %v to int: %v\n", value, err)
				}
				controlIdField := control.Field(valueTypeIndex)
				if controlIdField.IsValid() {
					controlIdField.SetInt(int64(val))
				}
			} else if key == "editType" {
				val, err := strconv.Atoi(value)
				if err != nil {
					return fmt.Errorf("editType error converting %v to int: %v\n", value, err)
				}
				controlIdField := control.Field(editTypeIndex)
				if controlIdField.IsValid() {
					controlIdField.SetInt(int64(val))
				}
			} else if key == "default" {
				valueField := control.Field(valueIndex)
				if valueField.IsValid() {
					valueField.SetString(value)
				}
			}
		}
	}
	return nil
}

// setControlValue 把srcVal的值赋值到descVal
func setControlValue(descVal reflect.Value, srcVal reflect.Value, structField reflect.StructField) error {
	if srcVal.IsZero() {
		return nil
	}
	// 把structField这个属性的赋值到ControlDTO的value属性上
	// string类型会setString
	// Array和Slice会通过json格式化的方式赋值成json字符串
	// 其他情况需要需要实现fmt.Stringer接口返回string来赋值
	switch srcVal.Kind() {
	case reflect.String:
		descVal.SetString(srcVal.String())
	case reflect.Array, reflect.Slice:
		bytes, err := jsoniter.Marshal(srcVal.Interface())
		if err != nil {
			return fmt.Errorf("json.Marshal %v to int: %v\n", srcVal.Interface(), err)
		}
		descVal.SetString(string(bytes))
	default:
		// 判断该类型是否实现了fmt.Stringer接口, 实现了则调用 string()赋值
		// 注意: 目前好像不支持指针接收实现fmt.Stringer接口, 不要用指针接收实现 fmt.Stringer
		stringerImpl := reflect.TypeOf((*fmt.Stringer)(nil)).Elem()
		if structField.Type.Implements(stringerImpl) {
			s := srcVal.Interface().(fmt.Stringer)
			descVal.SetString(s.String())
		} else {
			return fmt.Errorf("unsupported kind %v", srcVal.Kind())
		}
	}
	return nil
}
