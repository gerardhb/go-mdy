package mdy

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type (
	App struct {
		AppId     string    `json:"appId"`     // "应用id"
		Color     string    `json:"color"`     // "图标颜色"
		Desc      string    `json:"desc"`      // "应用描述"
		IconUrl   string    `json:"iconUrl"`   // "图标地址"
		Name      string    `json:"name"`      // "应用名称"
		ProjectId string    `json:"projectId"` // "网络id"
		Sections  []Section `json:"sections"`
	}

	Section struct {
		Name      string `json:"name"`      // "分组名称"
		SectionId string `json:"sectionId"` // "应用分组id"
		Items     []struct {
			Alias   string `json:"alias"`   // "工作表别名"
			IconUrl string `json:"iconUrl"` // "应用项图标地址"
			Id      string `json:"id"`      // "分组下应用项id"
			Name    string `json:"name"`    // "应用项名称"
			Status  int    `json:"status"`
			Type    int    `json:"type"`
		} `json:"items"`
	}
)

// AppRequest 代表明道云应用信息Request
type AppRequest struct {
	mdy *Client
	Req *resty.Request
}

// Get 获取应用信息 GET
func (r *AppRequest) Get() (*App, error) {
	response, err := r.Req.SetResult(&Response[App]{}).
		SetQueryParam("appKey", r.mdy.appKey).
		SetQueryParam("sign", r.mdy.sign).
		Get(AppGetURL)
	if err != nil {
		return nil, err
	}

	result := response.Result().(*Response[App])

	return &result.Data, nil
}

// GetSignature 明道云签名
func GetSignature(appKey, secretKey string, ti int64) string {
	//sha256
	key := fmt.Sprintf("AppKey=%v&SecretKey=%v&Timestamp=%v", appKey, secretKey, ti)
	h := sha256.New()
	h.Write([]byte(key))
	shaKey := h.Sum(nil)

	//hex 转string
	hexKey := hex.EncodeToString(shaKey)
	fmt.Println(hexKey)

	//base64
	return base64.StdEncoding.EncodeToString([]byte(hexKey))
}
